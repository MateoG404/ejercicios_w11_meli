package main

import (
	"ejercicio1/app/internal/application"
	"os"

	"github.com/go-sql-driver/mysql"
)

func main() {

	// Use the server to execute queries

	Passwd := os.Getenv("PASSWORD")
	mysqlCfg := mysql.Config{User: "root",
		Passwd:               Passwd,
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "melisprint",
		ParseTime:            true,
		AllowNativePasswords: true,
	}
	cfg := application.ServerChiConfig{
		Address:  ":8080",
		MySqlDSN: mysqlCfg.FormatDSN(),
	}

	chiserver := application.NewServerChi(cfg)

	chiserver.Start()

}
