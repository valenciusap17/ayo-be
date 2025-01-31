package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type ClaimSpec struct {
	userId string
	expirationTime *time.Time
}

type Claims struct {
	jwt.RegisteredClaims
}