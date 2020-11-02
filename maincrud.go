package main
 
import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)
 
type Post struct {
    id   int
    Name string
    Text string
}



db, e := sql.Open("mysql", "root:Akash@8888@tcp(127.0.0.1:3306)/post")
ErrorCheck(e)

defer db.Close()
PingDB(db)

func PingDB(db *sql.DB) {
    err := db.Ping()
    ErrorCheck(err)
}



func ErrorCheck(err error) {
    if err != nil {
        panic(err.Error())
    }
}