package database

import (
	"context"
	"database/sql"
	"github.com/mattrighetti/fdm-repository-backend/config"
	"log"
	"time"
)

func queryWithContext(query string, handler func(rows *sql.Rows) error, args ...interface{}) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()
	stmt, err := config.Db.PrepareContext(ctx, query)
	if err != nil {
		log.Fatalf("encountered error while querying database: %v", err)
	}
	defer stmt.Close()
	res, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		log.Fatalf("encountered error while querying database: %v", err)
	}
	return handler(res)
}