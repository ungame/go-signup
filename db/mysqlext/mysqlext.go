package mysqlext

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

const (
	driverName      = "mysql"
	maxOpenConns    = 25
	maxIdleConns    = 25
	maxConnLifetime = time.Minute * 5
)

func New(ctx context.Context, cfg MySQLConfig) *sql.DB {
	conn, err := sql.Open(driverName, cfg.String())
	if err != nil {
		log.Fatalln("unable to open connection with mysql:", err)
	}
	err = conn.PingContext(ctx)
	if err != nil {
		log.Fatalln("unable to ping mysql:", err)
	}
	conn.SetMaxOpenConns(maxOpenConns)
	conn.SetMaxIdleConns(maxIdleConns)
	conn.SetConnMaxLifetime(maxConnLifetime)
	return conn
}
