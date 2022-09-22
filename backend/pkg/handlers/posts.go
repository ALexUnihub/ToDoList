package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"toDoList/pkg/posts"
	// "toDoList/pkg/posts"
)

type PostHandler struct {
	PostsRepo posts.PostsRepo
}

func (h *PostHandler) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	dataHTML, err := os.Open("../../static/build/index.html")
	if err != nil {
		http.Error(w, `Template errror`, http.StatusInternalServerError)
		return
	}

	byteValue, err := ioutil.ReadAll(dataHTML)
	if err != nil {
		http.Error(w, `Parsing html error`, http.StatusInternalServerError)
		return
	}

	_, err = w.Write(byteValue)
	if err != nil {
		log.Println("err in pckg handlers, Index", err.Error())
		return
	}
}

func (h *PostHandler) SendAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.PostsRepo.GetAllPosts()
	if err != nil {
		log.Println("package handlers, SendAllPosts, GetAllPosts, err: ", err.Error())
		return
	}

	byteValue, err := json.Marshal(posts)
	if err != nil {
		log.Println("package handlers, SendAllPosts, Marshal, err: ", err.Error())
		return
	}

	_, err = w.Write(byteValue)
	if err != nil {
		log.Println("package handlers, SendAllPosts, Write, err: ", err.Error())
		return
	}
}

func (h *PostHandler) AddPost(w http.ResponseWriter, r *http.Request) {
	newPost, err := h.PostsRepo.AddPostRepo(r)
	if err != nil {
		log.Println("package handlers, AddPost, AddPostRepo, err: ", err.Error())
		return
	}

	byteValue, err := json.Marshal(newPost)
	if err != nil {
		log.Println("package handlers, AddPost, Marshal, err: ", err.Error())
		return
	}

	_, err = w.Write(byteValue)
	if err != nil {
		log.Println("package handlers, AddPost, Write, err: ", err.Error())
		return
	}
}
func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	postID := r.URL.Path
	postID = postID[len("/api/post/"):]

	ID, err := strconv.Atoi(postID)
	if err != nil {
		log.Println("package handlers, DeletePost, Atoi, err: ", err.Error())
		return
	}

	err = h.PostsRepo.DeletePostFromRepo(ID)
	if err != nil {
		log.Println("package handlers, DeletePost, DeletePostFromRepo, err: ", err.Error())
		return
	}
}

func (h *PostHandler) ChangePost(w http.ResponseWriter, r *http.Request) {
	postID := r.URL.Path
	postID = postID[len("/api/post/"):]

	ID, err := strconv.Atoi(postID)
	if err != nil {
		log.Println("package handlers, ChangePost, Atoi, err: ", err.Error())
		return
	}

	_, err = h.PostsRepo.ChangePostInRepo(ID, r)
	if err != nil {
		log.Println("package handlers, ChangePost, ChangePostInRepo, err: ", err.Error())
		return
	}
}
