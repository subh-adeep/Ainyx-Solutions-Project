package service

import (
	"context"
	db "project/db/sqlc"
	"project/internal/models"
	"project/internal/repository"
	"project/pkg/util"
	"time"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, req models.CreateUserRequest) (models.UserResponse, error) {
	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return models.UserResponse{}, err
	}

	arg := db.CreateUserParams{
		Name: req.Name,
		Dob:  dob,
	}

	user, err := s.repo.CreateUser(ctx, arg)
	if err != nil {
		return models.UserResponse{}, err
	}

	return s.mapToResponse(user), nil
}

func (s *UserService) GetUserByID(ctx context.Context, id int32) (models.UserResponse, error) {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return models.UserResponse{}, err
	}
	return s.mapToResponse(user), nil
}

func (s *UserService) ListUsers(ctx context.Context, page, limit int) ([]models.UserResponse, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit

	arg := db.ListUsersParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	users, err := s.repo.ListUsers(ctx, arg)
	if err != nil {
		return nil, err
	}

	resp := make([]models.UserResponse, len(users))
	for i, u := range users {
		resp[i] = s.mapToResponse(u)
	}
	return resp, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id int32, req models.UpdateUserRequest) (models.UserResponse, error) {
	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return models.UserResponse{}, err
	}

	arg := db.UpdateUserParams{
		ID:   id,
		Name: req.Name,
		Dob:  dob,
	}

	user, err := s.repo.UpdateUser(ctx, arg)
	if err != nil {
		return models.UserResponse{}, err
	}
	return s.mapToResponse(user), nil
}

func (s *UserService) DeleteUser(ctx context.Context, id int32) error {
	return s.repo.DeleteUser(ctx, id)
}

func (s *UserService) mapToResponse(u db.User) models.UserResponse {
	age := util.CalculateAge(u.Dob, time.Now())

	return models.UserResponse{
		ID:   u.ID,
		Name: u.Name,
		DOB:  u.Dob.Format("2006-01-02"),
		Age:  age,
	}
}
