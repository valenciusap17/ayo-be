package account

import (
	"ayo/cmd/config"
	"ayo/internal/token"
	"ayo/utils/errors"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type AccountDBAccessor interface {
	writeAccount(ctx context.Context, account *Account) error
    findAccountByEmail(ctx context.Context, email string) (*Account, error)

}

type AccountService struct {
    AccountDBAccessor
    *token.TokenService
    config *config.Application
}

func (s *AccountService) Register(ctx context.Context, spec AuthenticationSpec) (*Account, *errors.AppError) {
    passwordBytes := []byte(spec.Password)
    hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, 10)
    if err != nil {
        return nil, errors.UnauthorizedError(err.Error())
    }
    
    account := &Account{
        ID: uuid.NewString(),
        Email: spec.Email,
        Password: string(hashedPassword),
        CreatedDate: time.Now(),
        ModifiedDate: time.Now(),
    }
    
    if err = s.writeAccount(ctx, account); err != nil {
        return nil , errors.InternalServerError(err.Error())
    }

    return account, nil
}

func (s *AccountService) Login(ctx context.Context, spec AuthenticationSpec) (*string, *errors.AppError) {
    existingAccount, err := s.AccountDBAccessor.findAccountByEmail(ctx, spec.Email)
    if err != nil {
        return nil, errors.NotFoundError("Account not found")
    }

    if err = bcrypt.CompareHashAndPassword([]byte(existingAccount.Password), []byte(spec.Password)); err != nil {
        return nil, errors.UnauthorizedError("Incorrect password")
    }

    
    token, err := s.TokenService.GenerateToken(existingAccount.ID)
    if err != nil {
        fmt.Println(err.Error())
        return nil, errors.InternalServerError("Token creation error")
    }
    
    return token, nil
}

func NewAccountService(ctx context.Context, db *sqlx.DB, cfg *config.Application, tokenSvc *token.TokenService) *AccountService {
    return &AccountService{
        AccountDBAccessor: newPostgresAccountAccessor(db),
        config: cfg,
        TokenService: tokenSvc,
    }
}