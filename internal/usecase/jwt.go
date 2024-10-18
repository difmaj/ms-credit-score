package usecase

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/difmaj/ms-credit-score/internal/domain"
	"github.com/difmaj/ms-credit-score/internal/dto"
	"github.com/difmaj/ms-credit-score/internal/pkg/config"
	"github.com/google/uuid"
)

func (uc *Usecase) jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	signKey := []byte(config.Env.JWTGlobalKey)

	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	claims, ok := token.Claims.(*domain.Claims)
	if !ok {
		return nil, &dto.APIError{
			Status:  http.StatusUnauthorized,
			Message: "invalid claims",
		}
	}

	signKey, _ = uc.jwtGetSignKey(claims.User.ID.String(), claims.Ref.String())
	if signKey == nil {
		return nil, &dto.APIError{
			Status:  http.StatusUnauthorized,
			Message: "invalid claims",
		}
	}
	return signKey, nil
}

// ParseJWT parses a JWT token.
func (uc *Usecase) ParseJWT(token string) (*jwt.Token, error) {
	claims := &domain.Claims{}
	token = strings.TrimPrefix(token, "Bearer ")

	tokenParsed, err := jwt.ParseWithClaims(token, claims, uc.jwtKeyFunc)
	if err != nil {
		return nil, err
	}
	return tokenParsed, nil
}

// ClaimsJWT returns the claims of a JWT token.
func (uc *Usecase) ClaimsJWT(token string) (*domain.Claims, error) {
	tokenParsed, err := uc.ParseJWT(token)
	if err != nil {
		return nil, err
	}

	if claims, ok := tokenParsed.Claims.(*domain.Claims); ok {
		claims.Token = token
		return claims, nil
	}

	return nil, &dto.APIError{
		Status:  http.StatusUnauthorized,
		Message: "invalid claims",
	}
}

// GenerateToken generates a JWT token.
func (uc *Usecase) GenerateToken(user *domain.User, privileges domain.Privileges) (*domain.Claims, error) {
	refreshToken, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	signKey, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	exp := time.Now().Add(24 * time.Hour).Unix()

	claims := &domain.Claims{
		User:        *user,
		IssuedAt:    time.Now().Unix(),
		Subject:     uuid.NewString(),
		Audience:    config.Env.JWTAud,
		Issuer:      config.Env.JWTIss,
		Permissions: privileges,
		ExpiresAt:   exp,
		Ref:         refreshToken,
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(signKey.String()))
	if err != nil {
		return nil, err
	}
	claims.Token = token

	refreshExp := 86400 * config.Env.JWTRefreshDaysExp
	err = uc.jwtSaveSignKey(user.ID.String(), refreshToken.String(), signKey.String(), refreshExp)
	if err != nil {
		return nil, err
	}
	return claims, nil
}
