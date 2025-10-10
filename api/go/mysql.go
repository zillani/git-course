package main

import (
"database/sql"
"fmt"
"log"

_ "github.com/go-sql-driver/mysql"

)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/college"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to connect:", err)
	}
	rows, err := db.Query("SELECT id, name, branch FROM students")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name, branch string
		if err := rows.Scan(&id, &name, &branch); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s, Branch: %s\n", id, name, branch)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MySQL successfully!")
}

