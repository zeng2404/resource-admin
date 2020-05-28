package repository

type TagRepository interface {

}

func NewTagRepository() TagRepository {
	return &tagRepository{}
}

type tagRepository struct {

}
