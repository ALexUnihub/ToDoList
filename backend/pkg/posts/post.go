package posts

import "net/http"

type Post struct {
	ID     int    `json:"ID"`
	Title  string `json:"Title"`
	Text   string `json:"Text"`
	IsDone bool   `json:"IsDone"`
}

type PostsRepo interface {
	GetAllPosts() ([]*Post, error)
	AddPostRepo(r *http.Request) (*Post, error)
	DeletePostFromRepo(postID int) error
	ChangePostInRepo(postID int, r *http.Request) (*Post, error)
}
