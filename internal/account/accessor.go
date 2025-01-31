package account

import (
	"context"

	"github.com/jmoiron/sqlx"
)


type postgresAccountAccessor struct {
    db *sqlx.DB
}

const (
    insertAccountQuery = `
        INSERT INTO account 
            (id, email, password, username, fullname, phone_number, created_date, modified_date)
        VALUES
            (:id, :email, :password, :username, :fullname, :phone_number, :created_date, :modified_date)
    `

    findAccountByEmailQuery = `
        SELECT * FROM account WHERE email = $1
    `
)

func (p *postgresAccountAccessor) writeAccount(ctx context.Context, user *Account) error {
    if _, err := p.db.NamedExec(insertAccountQuery, user); err != nil {
        return err
    }
    return nil
}

func (p *postgresAccountAccessor) findAccountByEmail(ctx context.Context, email string) (*Account, error) {
    account := &Account{}
    if err := p.db.Get(account, findAccountByEmailQuery, email); err != nil {
        return nil, err
    }
    
    return account, nil
}

func newPostgresAccountAccessor(db *sqlx.DB) *postgresAccountAccessor {
    return &postgresAccountAccessor{
        db: db,
    }
}