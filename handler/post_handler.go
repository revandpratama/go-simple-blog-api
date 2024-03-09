package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/revandpratama/go-simple-blog-api/dto"
	"github.com/revandpratama/go-simple-blog-api/errorhandler"
	"github.com/revandpratama/go-simple-blog-api/helper"
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
	posts, err := h.postService.GetAll()
	if err != nil {
		errorhandler.HandleError(g, err)
		return
	}

	res := helper.Response(dto.ResponseParam{
		StatusCode: http.StatusOK,
		Message:    "Success get all posts",
		Data:       posts,
	})

	g.JSON(http.StatusOK, res)
}
