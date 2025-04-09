package connections

import (
	"go-gym-auth-service/internal/app/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Connections struct {
	DB *sqlx.DB
}

func New(cfg *config.Config) (*Connections, error) {
	db, err := sqlx.Connect("postgres", cfg.DB.DSN)
	if err != nil {
		return nil, err
	}
	return &Connections{DB: db}, nil
}

func (c *Connections) Close() {
	if c.DB != nil {
		_ = c.DB.Close()
	}
}
