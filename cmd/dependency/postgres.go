package dependency

import (
	"ayo/cmd/config"
	"ayo/internal/common/database"

	"github.com/jmoiron/sqlx"
)

func NewPostgreSQL(config config.Postgres) *sqlx.DB {
    return database.NewDBConnection(
        config.User, 
        config.Password, 
        config.Name, 
        config.Port, 
        config.Host,
    )
}