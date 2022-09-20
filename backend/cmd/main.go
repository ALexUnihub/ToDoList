package main

import (
	"log"
	"net/http"
	"toDoList/pkg/handlers"
	"toDoList/pkg/posts"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	postsRepo := posts.NewMemoryRepo()
	postsRepo.Data, _ = posts.AddStaticPosts()
	postsRepo.ID = 3

	postHandler := handlers.PostHandler{
		PostsRepo: postsRepo,
	}

	r.HandleFunc("/", postHandler.Index).Methods("GET")
	r.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("../../static/"))))

	r.HandleFunc("/api/posts", postHandler.SendAllPosts).Methods("GET")
	r.HandleFunc("/api/post", postHandler.AddPost).Methods("POST")
	r.HandleFunc("/api/post/{POST_ID}", postHandler.DeletePost).Methods("DELETE")
	r.HandleFunc("/api/post/{POST_ID}", postHandler.ChangePost).Methods("POST")

	addr := ":8080"
	log.Println("server starting on addr", addr)
	err := http.ListenAndServe(addr, r)

	if err != nil {
		log.Println("err: ", err.Error())
	}
}
