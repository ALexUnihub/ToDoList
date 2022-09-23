package posts

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ItemMemoryRepository struct {
	ID        int
	Data      []*Post
	DataDB    *mgo.Collection
	SessionDB *mgo.Session
}

func NewMemoryRepo(sess *mgo.Session) *ItemMemoryRepository {
	collection := sess.DB("posts").C("postRepo")

	AddStaticPostsDB(collection, sess)

	return &ItemMemoryRepository{
		ID:        0,
		DataDB:    collection,
		SessionDB: sess,
	}
}

type MyPostForm struct {
	Title  string `json:"Title"`
	Text   string `json:"Text"`
	IsDone bool   `json:"IsDone"`
}

func (repo *ItemMemoryRepository) GetAllPosts() ([]*Post, error) {

	posts := []*Post{}
	err := repo.DataDB.Find(bson.M{}).All(&posts)
	if err != nil {
		log.Println("package posts, GetAllPosts, Find, err: ", err.Error())
		return nil, err
	}

	log.Println("Get all posts")
	return posts, nil
}

func (repo *ItemMemoryRepository) AddPostRepo(r *http.Request) (*Post, error) {
	body, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()

	postData := &MyPostForm{}
	err := json.Unmarshal(body, postData)

	if err != nil {
		log.Println("package posts, AddPost, Unmarshal, err: ", err.Error())
		return nil, err
	}

	newPost := &Post{
		ID:     repo.ID,
		Title:  postData.Title,
		Text:   postData.Text,
		IsDone: false,
	}

	// insert in MongoDB
	newPost.IDBson = bson.NewObjectId()
	err = repo.DataDB.Insert(&newPost)
	if err != nil {
		log.Println("package posts, AddPostRepo, Insert, err: ", err.Error())
		return nil, err
	}
	//

	repo.ID++
	log.Println("new post added to repo, post", newPost)

	return newPost, nil
}

func (repo *ItemMemoryRepository) DeletePostFromRepo(postID int) error {
	err := repo.DataDB.Remove(bson.M{"ID": postID})
	if err != nil {
		log.Println("package posts, DeletePostFromRepo, Remove, err: ", err.Error())
		return err
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
		log.Println("package posts, ChangePostInRepo, Unmarshal, err: ", err.Error())
		return nil, err
	}

	changedPost := &Post{
		ID:     postID,
		Title:  postData.Title,
		Text:   postData.Text,
		IsDone: postData.IsDone,
	}

	postInRepo := &Post{}
	err = repo.DataDB.Find(bson.M{"ID": postID}).One(&postInRepo)
	if err != nil {
		log.Println("package posts, ChangePostInRepo, Find, err: ", err.Error())
		return nil, err
	}

	changedPost.IDBson = postInRepo.IDBson

	err = repo.DataDB.Update(bson.M{"ID": postID}, &changedPost)
	if err != nil {
		log.Println("package posts, ChangePostInRepo, Update, err: ", err.Error())
		return nil, err
	}

	log.Printf("Post changed in repo\tbefore: %v\tafter: %v", postInRepo, changedPost)
	return changedPost, nil
}

func AddStaticPostsDB(collection *mgo.Collection, sess *mgo.Session) ([]*Post, error) {
	post1 := &Post{
		ID:     0,
		Title:  "Title_1",
		Text:   "text_1_server",
		IsDone: true,
	}
	post2 := &Post{
		ID:     1,
		Title:  "Title_2",
		Text:   "text_2_server",
		IsDone: false,
	}
	post3 := &Post{
		ID:     2,
		Title:  "Title_3",
		Text:   "text_3_server",
		IsDone: true,
	}
	posts := []*Post{}
	posts = append(posts, post1, post2, post3)

	post1.IDBson = bson.NewObjectId()
	post2.IDBson = bson.NewObjectId()
	post3.IDBson = bson.NewObjectId()

	collection.Insert(&post1)
	collection.Insert(&post2)
	collection.Insert(&post3)

	return posts, nil
}
