package db_layer

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"strings"
)

func GetMysqlConnection() *sql.DB {
	pwds, err := ReadPasswordsAsStrings("passwords.txt")
	fmt.Println("password is : ", pwds)
	dsn := "root:"+pwds[0]+"@tcp(127.0.0.1:3306)/college"
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

func ReadPasswordsAsStrings(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var pwds []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		pwds = append(pwds, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return pwds, nil
}


