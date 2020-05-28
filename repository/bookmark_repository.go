package repository

type BookmarkRepository interface {

}

func NewBookmarkRepository() BookmarkRepository {
	return &bookmarkRepository{}
}

type bookmarkRepository struct {

}