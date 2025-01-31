package token

import (
	"ayo/cmd/config"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type jwtManager struct {
    config *config.Application
}

func (j *jwtManager) createToken(spec *ClaimSpec) (*string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
        RegisteredClaims: jwt.RegisteredClaims{
            ID: uuid.NewString(),
            ExpiresAt: &jwt.NumericDate{Time: *spec.expirationTime},
            Subject: spec.userId,
        },
    })

    secret := []byte(j.config.Token.Secret)

    tokenString, err := token.SignedString(secret)
    if err != nil {
        return nil, err
    }

    return &tokenString, nil
}

func (j *jwtManager) verifyToken(tokenString string) error {
	secret := []byte(j.config.Token.Secret)
    
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(
        tokenString, 
        claims, 
        func(t *jwt.Token) (interface{}, error) {
		    return secret, nil
        },
        jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
        jwt.WithIssuedAt(),
    )

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

func NewJwtManager(config *config.Application) *jwtManager {
    return &jwtManager{
        config: config,
    }
}