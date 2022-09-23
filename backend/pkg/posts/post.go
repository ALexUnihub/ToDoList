package posts

import (
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

type Post struct {
	IDBson bson.ObjectId `json:"-" bson:"_id"`
	ID     int           `json:"ID" bson:"ID"`
	Title  string        `json:"Title" bson:"Title"`
	Text   string        `json:"Text" bson:"Text"`
	IsDone bool          `json:"IsDone" bson:"IsDone"`
}

type PostsRepo interface {
	GetAllPosts() ([]*Post, error)
	AddPostRepo(r *http.Request) (*Post, error)
	DeletePostFromRepo(postID int) error
	ChangePostInRepo(postID int, r *http.Request) (*Post, error)
}
