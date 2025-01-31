package token

import (
	"ayo/cmd/config"
	"time"
)

type TokenManager interface {
	createToken(spec *ClaimSpec) (*string, error)
	verifyToken(tokenString string) error
}

type TokenService struct {
    TokenManager
}

func (t *TokenService) GenerateToken(userId string) (*string, error){
    expiresAt := &time.Time{}
    *expiresAt = time.Now().Add(time.Hour)
    
    spec := &ClaimSpec{
        userId: userId,
        expirationTime: expiresAt,
    }

    return t.TokenManager.createToken(spec)
}

func (t *TokenService) VerifyToken(tokenString string) error {
    return t.TokenManager.verifyToken(tokenString)
}

func NewTokenService(config *config.Application) *TokenService {
    return &TokenService{
        TokenManager: NewJwtManager(config),
    }
}