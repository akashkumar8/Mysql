package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	fmt.Println("Go Connect to mysql -akashkr")

	db, err := sql.Open("mysql", "root:Akash@8888@tcp(127.0.0.1:3306)/mysql")

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connected")

	defer db.Close()
}
