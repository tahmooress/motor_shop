package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/tahmooress/motor-shop/internal/entities/interfaces"

	_ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
	db *sql.DB
}

func New(_ context.Context) (interfaces.IDatabase, error) {
	dbname := os.Getenv("MYSQL_DBNAME")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASS")

	if dbname == "" || host == "" || port == "" || user == "" || pass == "" {
		return nil, fmt.Errorf("some envorirment variables are empty")
	}

	db, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, dbname),
	)

	if err != nil {
		return nil, fmt.Errorf("new() >> sql.Open() >> %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("new() >> db.Ping() >> %w", err)
	}

	return &Mysql{
		db: db,
	}, nil
}
