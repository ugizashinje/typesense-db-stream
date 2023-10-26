package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "db:db@tcp(127.0.0.1:55677)/db")
	if err != nil {
		panic(err)
	}

	fmt.Println("Hello, shiggy.")
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	os.Exit(-1)
}
