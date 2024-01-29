package main

import (
	"database/sql"
	"fmt"

	"os"

	"github.com/go-sql-driver/mysql"
)

func main() {
	// database
	// config
	Passwd := os.Getenv("PASSWORD")
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               Passwd,
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "melisprint",
		ParseTime:            true,
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Test connection using Ping
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database!")
}
