package repository

type BookmarkTagRepository interface {

}

type bookmarkTagRepository struct {

}

func NewBookmarkTagRepository() BookmarkTagRepository{
	return bookmarkTagRepository{}
}
