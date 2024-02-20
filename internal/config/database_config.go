package config

import (
	"context"
	"fmt"
	"log"

	"github.com/PNYwise/user-service/internal/domain"
	"github.com/jackc/pgx/v5"
)

var db *pgx.Conn

func getDatabaseConn(ctx context.Context, extConf *domain.ExtConf) *pgx.Conn {
	dbConfig := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		extConf.Database.Username,
		extConf.Database.Password,
		extConf.Database.Host,
		extConf.Database.Port,
		extConf.Database.Name,
	)
	connConfig, err := pgx.ParseConfig(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	newDB, err := pgx.ConnectConfig(ctx, connConfig)
	if err != nil {
		log.Fatal(err)
	}

	if err := newDB.Ping(ctx); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to Database")
	db = newDB
	return db
}
func DbConn(ctx context.Context, extConf *domain.ExtConf) *pgx.Conn {
	if db == nil {
		db = getDatabaseConn(ctx, extConf)
	}
	return db
}
