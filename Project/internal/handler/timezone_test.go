package handler_test

import (
	"context"
	db "project/db/sqlc"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/mock"
)

// Mock Repository
type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).(db.User), args.Error(1)
}
func (m *MockRepo) GetUserByID(ctx context.Context, id int32) (db.User, error) { return db.User{}, nil }
func (m *MockRepo) ListUsers(ctx context.Context, arg db.ListUsersParams) ([]db.User, error) {
	return nil, nil
}
func (m *MockRepo) UpdateUser(ctx context.Context, arg db.UpdateUserParams) (db.User, error) {
	return db.User{}, nil
}
func (m *MockRepo) DeleteUser(ctx context.Context, id int32) error { return nil }

func TestDOBValidation_Timezones(t *testing.T) {
	// We verify the validation logic in isolation to catch timezone issues.

	v := validator.New()
	v.RegisterValidation("dob_past", func(fl validator.FieldLevel) bool {
		dateStr := fl.Field().String()
		now := time.Now()
		dob, err := time.ParseInLocation("2006-01-02", dateStr, now.Location())
		if err != nil {
			return false
		}

		today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		return !dob.After(today)
	})

	// Test Case 1: Born "Today"
	// Using simple format
	now := time.Now()
	todayStr := now.Format("2006-01-02")

	s1 := struct {
		DOB string `validate:"dob_past"`
	}{DOB: todayStr}

	if err := v.Struct(s1); err != nil {
		t.Errorf("Born Today (%s) Validation Error: %v", todayStr, err)
	}

	// Test Case 2: Born "Tomorrow" (Should Fail)
	tomorrow := now.AddDate(0, 0, 1)
	tomorrowStr := tomorrow.Format("2006-01-02")
	s2 := struct {
		DOB string `validate:"dob_past"`
	}{DOB: tomorrowStr}

	if err := v.Struct(s2); err == nil {
		t.Errorf("Born Tomorrow (%s) should have failed validation", tomorrowStr)
	} else {
		t.Logf("Born Tomorrow (%s) correctly failed validation", tomorrowStr)
	}
}
