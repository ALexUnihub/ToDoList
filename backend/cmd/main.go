package main

import (
	_ "github.com/go-sql-driver/mysql"

	"log"
	"net/http"
	"toDoList/pkg/handlers"
	"toDoList/pkg/posts"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

func main() {
	// test MySQL
	// dsn := "root:love@tcp(localhost:3306)/golang?"
	// dsn += "charset=utf8"
	// dsn += "&interpolateParams=true"

	// db, err1 := sql.Open("mysql", dsn)

	// if err1 != nil {
	// 	log.Println(err1.Error())
	// }

	// db.SetMaxOpenConns(10)

	// err1 = db.Ping() // вот тут будет первое подключение к базе
	// if err1 != nil {
	// 	log.Println("BAD PING")
	// 	panic(err1)
	// }

	// mongo DB
	sessMongodb, err := mgo.Dial("localhost")
	// sessMongodb, err := mgo.Dial("mongodb://root:example@localhost:27017")
	if err != nil {
		log.Println("sess mongodb err")
		panic(err)
	}
	//

	r := mux.NewRouter()

	postsRepo := posts.NewMemoryRepo(sessMongodb)

	postHandler := handlers.PostHandler{
		PostsRepo: postsRepo,
	}

	r.HandleFunc("/", postHandler.Index).Methods("GET")
	r.PathPrefix("/static/js/").Handler(
		http.StripPrefix("/static/js/",
			http.FileServer(http.Dir("../../static/build/static/js"))))
	r.PathPrefix("/static/css/").Handler(
		http.StripPrefix("/static/css/",
			http.FileServer(http.Dir("../../static/build/static/css"))))
	r.PathPrefix("/manifest.json").Handler(
		http.StripPrefix("/manifest.json",
			http.FileServer(http.Dir("../../static/build"))))

	r.HandleFunc("/api/posts", postHandler.SendAllPosts).Methods("GET")
	r.HandleFunc("/api/post", postHandler.AddPost).Methods("POST")
	r.HandleFunc("/api/post/{POST_ID}", postHandler.DeletePost).Methods("DELETE")
	r.HandleFunc("/api/post/{POST_ID}", postHandler.ChangePost).Methods("POST")

	addr := ":8080"
	log.Println("server starting on addr", addr)
	err = http.ListenAndServe(addr, r)

	if err != nil {
		log.Println("err: ", err.Error())
	}
}
