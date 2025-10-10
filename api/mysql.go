package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zillani/git-course/db_layer"
	"log"
)

func main() {
	db := db_layer.GetMysqlConnection()
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
	defer db.Close()
}

