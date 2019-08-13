package postgres

import (
	"fmt"

	"../../internal/blog"
	"../../internal/comment"

	// "../../internal/thread"
	_ "github.com/lib/pq"

	// "../../internal/user"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

type SQLRepository interface {
	comment.Repository
	blog.Repository
	// user.Repository
	// thread.Repository
}

type Config struct {
	Host     string
	Name     string
	Port     uint
	User     string
	Password string
}

func NewDB(cfg Config) (SQLRepository, error) {
	connString := fmt.Sprintf("host=%s, port=%d user=%s dbname=%s password=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Name, cfg.Password)
	db, err := sqlx.Open("postgres", connString)
	return &repository{db}, err
}
