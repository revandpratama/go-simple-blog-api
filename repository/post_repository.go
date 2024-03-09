package repository

import (
	"github.com/revandpratama/go-simple-blog-api/entity"
	"github.com/revandpratama/go-simple-blog-api/errorhandler"
	"gorm.io/gorm"
)

type PostRepository interface {
	GetAll() (*[]entity.Post, error)
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *postRepository {
	return &postRepository{
		db: db,
	}
}

func (r *postRepository) GetAll() (*[]entity.Post, error) {
	var post []entity.Post

	if err := r.db.Find(&post); err != nil {
		return nil, &errorhandler.NotFoundError{Message:  "Post not found"}
	}

	return &post, nil
}
