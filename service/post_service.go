package service

import "github.com/revandpratama/go-simple-blog-api/repository"

type PostService interface {
}

type postService struct {
	postRepository repository.PostRepository
}

func NewPostService(r repository.PostRepository) *postService {
	return &postService{
		postRepository: r,
	}
}
