package models



// CreateUserRequest represents the request body for creating a user.
type CreateUserRequest struct {
	Name string `json:"name" validate:"required,min=1"`
	DOB  string `json:"dob" validate:"required,dob_past"` // Format: 2006-01-02
}

// UpdateUserRequest represents the request body for updating a user.
type UpdateUserRequest struct {
	Name string `json:"name" validate:"required,min=1"`
	DOB  string `json:"dob" validate:"required,dob_past"` // Format: 2006-01-02
}

// UserResponse represents the response body for a user.
type UserResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	DOB  string `json:"dob"`
	Age  int    `json:"age"`
}
