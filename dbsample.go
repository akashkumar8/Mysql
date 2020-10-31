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

	rows, err := db.Query("SELECT * FROM student")

	checkErr(err)

	for rows.Next() {

		var studentid int
		var name string
		var roll int
		var studentcol string
		err = rows.Scan(&studentid, &name, &roll, &studentcol)
		checkErr(err)
		fmt.Println(studentid)
		fmt.Println(name)
		fmt.Println(roll)
		fmt.Println(studentcol)
	}

	defer db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
