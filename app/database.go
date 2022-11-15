package app

import (
	"database/sql"
	"fmt"
	"github.com/aziemp66/go-restful-api/helper"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func NewDB(
	DB_USER string,
	DB_PASSWORD string,
	DB_HOST string,
	DB_PORT string,
	DB_NAME string,
) *sql.DB {
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=true",
			DB_USER,
			DB_PASSWORD,
			DB_HOST,
			DB_PORT,
			DB_NAME,
		),
	)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
