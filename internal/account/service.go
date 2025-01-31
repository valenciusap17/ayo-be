package account

import (
	"ayo/utils/errors"
	"context"
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

func (s *AccountService) Login(ctx context.Context, spec AuthenticationSpec) *errors.AppError {
    existingAccount, err := s.AccountDBAccessor.findAccountByEmail(ctx, spec.Email)
    if err != nil {
        return errors.NotFoundError("Account not found")
    }

    if err = bcrypt.CompareHashAndPassword([]byte(existingAccount.Password), []byte(spec.Password)); err != nil {
        return errors.UnauthorizedError("Incorrect password")
    }

    return nil
}

func NewAccountService(ctx context.Context, db *sqlx.DB) *AccountService {
    return &AccountService{
        AccountDBAccessor: newPostgresAccountAccessor(db),
    }
}