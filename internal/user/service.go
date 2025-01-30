package user

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type UserDBAccessor interface {
	writeUser(ctx context.Context, user *User) error
}

type UserService struct {
    UserDBAccessor
}

func (s *UserService) Createuser(ctx context.Context, spec CreateUserSpec) (*User, error) {
    passwordBytes := []byte(spec.Password)
    hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, 10)
    if err != nil {
        return nil, err  
    }
    
    newUser := &User{
        ID: uuid.NewString(),
        Email: spec.Email,
        Password: string(hashedPassword),
        CreatedDate: time.Now(),
        ModifiedDate: time.Now(),
    }
    
    if err = s.writeUser(ctx, newUser); err != nil {
        return nil , err
    }

    return newUser, nil
}

func NewUserService(ctx context.Context, db *sqlx.DB) *UserService {
    return &UserService{
        UserDBAccessor: newPostgresUserAccessor(db),
    }
}