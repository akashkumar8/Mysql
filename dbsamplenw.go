package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	fmt.Println("Go Connect to mysql -akashkr")

	db, err := sql.Open("mysql", "root:Akash@8888@tcp(127.0.0.1:3306)/akash")

	if err != nil {
		panic(err.Error())
	}

	rows, err := db.Query("SELECT * FROM students")

	checkErr(err)

	for rows.Next() {

		var idstudents int
		var name string 
		var roll int
		err = rows.Scan(&idstudents, &name, &roll)
		checkErr(err)
		fmt.Println(idstudents)
		fmt.Println(name)
		fmt.Println(roll)
		
	}

	defer db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}