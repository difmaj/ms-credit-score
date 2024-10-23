package domain

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

// Privileges represents the permissions of a user.
type Privileges = map[string][]string

// Claims represents the claims of a JWT token.
type Claims struct {
	User        User       `json:"user"`
	Permissions Privileges `json:"permissions"`
	Token       string     `json:"token,omitempty"`
	Subject     string     `json:"sub"`
	Audience    string     `json:"aud"`
	ExpiresAt   int64      `json:"exp"`
	IssuedAt    int64      `json:"iat"`
	Issuer      string     `json:"iss"`
	Ref         uuid.UUID  `json:"ref"`
}

// Valid validates the claims of a JWT token.
func (cl *Claims) Valid() error {
	_, userIDErr := uuid.Parse(cl.Subject)

	expired := time.Now().Unix() >= cl.ExpiresAt
	if cl.Audience != os.Getenv("JWT_AUD") ||
		cl.Issuer != os.Getenv("JWT_ISS") ||
		userIDErr != nil ||
		expired {
		return errors.New(http.StatusText(http.StatusUnauthorized))
	}
	return nil
}
