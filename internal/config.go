package internal

import (
	"context"
	"database/sql"

	"github.com/Aritra640/docker-compose-eg/Database/db"
)

type Config struct {
	DB       *sql.DB
	QueryObj *db.Queries
	CTX      context.Context
}

var App *Config
