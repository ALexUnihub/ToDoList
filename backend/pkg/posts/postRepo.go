package posts

type ItemMemoryRepository struct {
	Length int
	Data   []*Post
}

func NewMemoryRepo() *ItemMemoryRepository {
	return &ItemMemoryRepository{
		Length: 0,
	}
}

func (repo *ItemMemoryRepository) GetAllPosts() ([]*Post, error) {
	return repo.Data, nil
}

func (repo *ItemMemoryRepository) AddPost(newPost *Post) (*Post, error) {
	repo.Data = append(repo.Data, newPost)

	return
}

func AddStaticPosts() ([]*Post, error) {
	post1 := &Post{
		Number: 1,
		Title:  "Title_1",
		Text:   "text_1",
		IsDone: true,
	}
	post2 := &Post{
		Number: 2,
		Title:  "Title_2",
		Text:   "text_2",
		IsDone: false,
	}
	post3 := &Post{
		Number: 3,
		Title:  "Title_3",
		Text:   "text_3",
		IsDone: true,
	}
	posts := []*Post{}
	posts = append(posts, post1, post2, post3)

	return posts, nil
}
