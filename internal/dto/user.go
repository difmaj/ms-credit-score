package dto

import "github.com/google/uuid"

// LoginHTTPInput represents the input of the login endpoint.
type LoginHTTPInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// LoginHTTPOutput represents the output of the login endpoint.
type LoginHTTPOutput struct {
	ID           uuid.UUID `json:"id"`
	AccessToken  string    `json:"access_token"`
	RefreshToken uuid.UUID `json:"refresh_token"`
}
