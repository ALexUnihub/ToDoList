package main

import (
	"fmt"
	"log"
	"net/http"
	"toDoList/pkg/handlers"
	"toDoList/pkg/posts"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello")

	r := mux.NewRouter()

	postsRepo := posts.NewMemoryRepo()
	postsRepo.Data, _ = posts.AddStaticPosts()

	postHandler := handlers.PostHandler{
		PostsRepo: postsRepo,
	}

	r.HandleFunc("/", postHandler.Index).Methods("GET")
	r.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("../../static/"))))

	r.HandleFunc("/GetAllPosts", postHandler.SendAllPosts).Methods("GET")
	r.HandleFunc("/AddPost", postHandler.SendAllPosts).Methods("GET")

	addr := ":8080"
	log.Println("server starting on addr", addr)
	err := http.ListenAndServe(addr, r)

	if err != nil {
		log.Println("err: ", err.Error())
	}
}
