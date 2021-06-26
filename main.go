package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

const (
	host     = "ec2-34-225-167-77.compute-1.amazonaws.com"
	port     = 5432
	user     = "bcpahssazerted"
	password = "49e06246ca80b646649e854d9b69bcaf21207ed3d8827d36ecf018e08094f0a4"
	dbname   = "dbbrrjvdi9ba20"
)

type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var posts []Post
	result, err := db.Query("SELECT id, title from posts")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var post Post
		err := result.Scan(&post.ID, &post.Title)
		if err != nil {
			panic(err.Error())
		}
		posts = append(posts, post)
	}
	json.NewEncoder(w).Encode(posts)
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=require",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Successfully connected!")

	router := mux.NewRouter()

	router.HandleFunc("/posts", getPosts).Methods("GET")

	http.ListenAndServe(":8080", router)

}
