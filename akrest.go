package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Post struct {
	ID    string `json:"id"`
	Title string `json:"name"`
	CITY  string `json:"city"`
}

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "root:Akash@8888@tcp(127.0.0.1:3306)/akash")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	router := mux.NewRouter()
	router.HandleFunc("/posts", getPosts).Methods("GET")
	//   router.HandleFunc("/posts", createPost).Methods("POST")
	//   router.HandleFunc("/posts/{id}", getPost).Methods("GET")
	//   router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	//   router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}
func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var posts []Post
	result, err := db.Query("SELECT id, name, city from state")
	fmt.Print(err, result)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var post Post
		err := result.Scan(&post.ID, &post.Title, &post.CITY)
		if err != nil {
			panic(err.Error())
		}
		posts = append(posts, post)
	}
	json.NewEncoder(w).Encode(posts)
}
