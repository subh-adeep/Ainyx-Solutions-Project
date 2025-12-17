package repository

import (
	"context"
	db "project/db/sqlc"
)

// UserRepository defines the interface for user data operations.
type UserRepository interface {
	CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error)
	GetUserByID(ctx context.Context, id int32) (db.User, error)
	ListUsers(ctx context.Context, arg db.ListUsersParams) ([]db.User, error)
	UpdateUser(ctx context.Context, arg db.UpdateUserParams) (db.User, error)
	DeleteUser(ctx context.Context, id int32) error
}

// SQLCUserRepository implements UserRepository using sqlc generated code.
type SQLCUserRepository struct {
	queries *db.Queries
}

func NewUserRepository(q *db.Queries) UserRepository {
	return &SQLCUserRepository{
		queries: q,
	}
}

func (r *SQLCUserRepository) CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error) {
	return r.queries.CreateUser(ctx, arg)
}

func (r *SQLCUserRepository) GetUserByID(ctx context.Context, id int32) (db.User, error) {
	return r.queries.GetUserByID(ctx, id)
}

func (r *SQLCUserRepository) ListUsers(ctx context.Context, arg db.ListUsersParams) ([]db.User, error) {
	return r.queries.ListUsers(ctx, arg)
}

func (r *SQLCUserRepository) UpdateUser(ctx context.Context, arg db.UpdateUserParams) (db.User, error) {
	return r.queries.UpdateUser(ctx, arg)
}

func (r *SQLCUserRepository) DeleteUser(ctx context.Context, id int32) error {
	return r.queries.DeleteUser(ctx, id)
}
