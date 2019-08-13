package blog

type Post struct {
	id      uint
	OwnerID uint
	UUID    string
	Title   string
	Body    string
}

type Info struct {
	OwnerID uint
	Title   string
	Body    string
}

func New(info Info) *Post {
	post := &Post{}
	post.OwnerID = info.OwnerID
	post.Title = info.Title
	post.Body = info.Body
	return post
}

type Repository interface {
	Service
}

func NewService(r Repository) Service {
	return storageService{r}
}

type storageService struct {
	Repository
}

type ID uint

type Fetcher interface {
	FetchPostByID(id ID) (*Post, error)
	FetchAllPosts(offset, limit uint) ([]*Post, error)
}

type Searcher interface {
	SearchPosts(name string) ([]*Post, error)
}

type Updater interface {
	UpdatePost(p *Post) (*Post, error)
}

type Remover interface {
	RemovePost(p *Post) error
}

type Creator interface {
	CreatePost(info *Info) (*Post, error)
}

type Service interface {
	Creator
	Fetcher
	Remover
	Updater
	Searcher
}
