package service

import (
	"github.com/revandpratama/go-simple-blog-api/dto"
	"github.com/revandpratama/go-simple-blog-api/repository"
)

type PostService interface {
	GetAll() (*[]dto.PostResponse, error)
}

type postService struct {
	postRepository repository.PostRepository
}

func NewPostService(r repository.PostRepository) *postService {
	return &postService{
		postRepository: r,
	}
}

func (s *postService) GetAll() (*[]dto.PostResponse, error) {

	var res []dto.PostResponse

	posts, err := s.postRepository.GetAll()
	if err != nil {
		return nil, err
	}

	for _, post := range *posts {
		res = append(res,  dto.PostResponse{
			ID : post.ID,
			Title:  post.Title,
			Excerpt: post.Excerpt,
			ImageURL:  post.ImageURL,
			Body: post.Body,
		})
	}

	return &res, nil
}
