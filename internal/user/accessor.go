package user

import (
	"context"

	"github.com/jmoiron/sqlx"
)


type postgresUserAccessor struct {
    db *sqlx.DB
}

const (
    insertUserQuery = `
        INSERT INTO user 
            (id, email, password, username, fullname, phone_number, modified_by, created_date, modified_date)
        VALUES
            (:id, :email, :password, :username, :fullname, :phone_number, :modified_by, :created_date, :modified_date)
    `
)

func (p *postgresUserAccessor) writeUser(ctx context.Context, user *User) error {
    if _, err := p.db.NamedExec(insertUserQuery, user); err != nil {
        return err
    }
    return nil
}


func newPostgresUserAccessor(db *sqlx.DB) *postgresUserAccessor {
    return &postgresUserAccessor{
        db: db,
    }
}