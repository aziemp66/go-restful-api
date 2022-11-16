package app

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/aziemp66/go-restful-api/helper"
	_ "github.com/go-sql-driver/mysql"
)

func NewDB(
	DB_USER string,
	DB_PASSWORD string,
	DB_HOST string,
	DB_PORT string,
	DB_NAME string,
) *sql.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		DB_USER,
		DB_PASSWORD,
		DB_HOST,
		DB_PORT,
		DB_NAME,
	)
	db, err := sql.Open("mysql", connectionString)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(60 * time.Minute)
	db.SetConnMaxLifetime(10 * time.Minute)

	return db
}
