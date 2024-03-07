package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/revandpratama/go-simple-blog-api/service"
)

// type PostHandler interface {
// 	GetAll(g *gin.Context) []entity.Post
// }

type postHandler struct {
	postService service.PostService
}

func NewPostHandler(s service.PostService) *postHandler {
	return &postHandler{
		postService: s,
	}
}

func (h *postHandler) GetAll(g *gin.Context) {
	// post, err := h.postService.GetAll()

	// g.JSON(http.StatusOK, res)
}
