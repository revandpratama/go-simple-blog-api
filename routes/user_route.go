package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/revandpratama/go-simple-blog-api/config"
	"github.com/revandpratama/go-simple-blog-api/handler"
	"github.com/revandpratama/go-simple-blog-api/repository"
	"github.com/revandpratama/go-simple-blog-api/service"
)

func UserRouter(api gin.RouterGroup) {
	userRepository := repository.NewUserRepository(config.DB)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	auth := api.Group("/auth")

	auth.POST("/login", userHandler.Login)
	auth.POST("/register", userHandler.Register)
}
