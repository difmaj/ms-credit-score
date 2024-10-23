package dto

import "github.com/google/uuid"

// LoginInput represents the input of the login endpoint.
type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// LoginOutput represents the output of the login endpoint.
type LoginOutput struct {
	ID           uuid.UUID `json:"id"`
	AccessToken  string    `json:"access_token"`
	RefreshToken uuid.UUID `json:"refresh_token"`
}
