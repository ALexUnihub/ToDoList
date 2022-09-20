package posts

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type ItemMemoryRepository struct {
	ID   int
	Data []*Post
}

func NewMemoryRepo() *ItemMemoryRepository {
	return &ItemMemoryRepository{
		ID: 0,
	}
}

type MyPostForm struct {
	Title  string `json:"Title"`
	Text   string `json:"Text"`
	IsDone bool   `json:"IsDone"`
}

func (repo *ItemMemoryRepository) GetAllPosts() ([]*Post, error) {
	log.Println("Get all posts")
	return repo.Data, nil
}

func (repo *ItemMemoryRepository) AddPostRepo(r *http.Request) (*Post, error) {
	body, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()

	postData := &MyPostForm{}
	err := json.Unmarshal(body, postData)

	if err != nil {
		log.Println("package handlers, AddPost, Unmarshal, err: ", err.Error())
		return nil, err
	}

	newPost := &Post{
		ID:     repo.ID,
		Title:  postData.Title,
		Text:   postData.Text,
		IsDone: false,
	}
	repo.Data = append(repo.Data, newPost)
	repo.ID++

	log.Println("new post added to repo, post", repo.Data[len(repo.Data)-1])

	return repo.Data[len(repo.Data)-1], nil
}

func (repo *ItemMemoryRepository) DeletePostFromRepo(postID int) error {

	for idx, item := range repo.Data {
		if item.ID == postID {
			copy(repo.Data[idx:], repo.Data[idx+1:])
			repo.Data[len(repo.Data)-1] = &Post{}
			repo.Data = repo.Data[:len(repo.Data)-1]
			break
		}
	}

	log.Println("Post deleted from repo, postID,", postID)

	return nil
}

func (repo *ItemMemoryRepository) ChangePostInRepo(postID int, r *http.Request) (*Post, error) {
	body, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()

	postData := &MyPostForm{}
	err := json.Unmarshal(body, postData)

	if err != nil {
		log.Println("package handlers, AddPost, Unmarshal, err: ", err.Error())
		return nil, err
	}

	for _, item := range repo.Data {
		if item.ID == postID {
			tmpPost := item

			item.Title = postData.Title
			item.Text = postData.Text
			item.IsDone = postData.IsDone

			log.Printf("post was changed: before:\t%v\tafter:\t%v", tmpPost, item)
			return item, nil
		}
	}
	return nil, nil
}

func AddStaticPosts() ([]*Post, error) {
	post1 := &Post{
		ID:     0,
		Title:  "Title_1",
		Text:   "text_1",
		IsDone: true,
	}
	post2 := &Post{
		ID:     1,
		Title:  "Title_2",
		Text:   "text_2",
		IsDone: false,
	}
	post3 := &Post{
		ID:     2,
		Title:  "Title_3",
		Text:   "text_3",
		IsDone: true,
	}
	posts := []*Post{}
	posts = append(posts, post1, post2, post3)

	return posts, nil
}
