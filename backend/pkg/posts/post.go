package posts

type Post struct {
	Number int    `json:"Number"`
	Title  string `json:"Title"`
	Text   string `json:"Text"`
	IsDone bool   `json:"IsDone"`
}

type PostsRepo interface {
	GetAllPosts() ([]*Post, error)
	AddPost(newPost *Post) (*Post, error)
}
