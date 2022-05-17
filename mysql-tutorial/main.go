package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `json:"name"`
}

func main() {
	fmt.Println("Go MySQL tutorial")

	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/testdb")

	if err != nil {
		fmt.Println("first error")
		panic(err.Error())
	}

	defer db.Close()

	// insert, err := db.Query("INSERT INTO users VALUES('habib')")

	// if err != nil {
	// 	fmt.Println("second error")
	// 	panic(err.Error())
	// }

	// defer insert.Close()

	// fmt.Println("Successfully inserted into user table")

	results, err := db.Query("SELECT name FROM users")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var user User

		err = results.Scan(&user.Name)

		if err != nil {
			panic(err.Error())
		}

		fmt.Println(user.Name)
	}

	fmt.Println("Successfully connected to mysql database")
}
