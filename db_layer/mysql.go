package db_layer

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func GetMysqlConnection() *sql.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/college"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MySQL successfully!")
	if err := db.Ping(); err != nil {
		log.Fatal("Failed to connect:", err)
	}
	return db
}

