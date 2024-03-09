package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/revandpratama/go-simple-blog-api/config"
	"github.com/revandpratama/go-simple-blog-api/handler"
	"github.com/revandpratama/go-simple-blog-api/repository"
	"github.com/revandpratama/go-simple-blog-api/service"
)

func PostRoutes(api *gin.RouterGroup) {
	postRepository := repository.NewPostRepository(config.DB)
	postService := service.NewPostService(postRepository)
	postHandler := handler.NewPostHandler(postService)
	post := api.Group("/posts")

	post.GET("/", postHandler.GetAll)
}
