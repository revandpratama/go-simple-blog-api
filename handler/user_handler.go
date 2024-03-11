package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/revandpratama/go-simple-blog-api/dto"
	"github.com/revandpratama/go-simple-blog-api/errorhandler"
	"github.com/revandpratama/go-simple-blog-api/service"
)

type authHandler struct {
	userService service.UserService
}

func NewAuthHandler(s service.UserService) *authHandler {
	return &authHandler{userService: s}
}

func (h *authHandler) Login(g *gin.Context) {
	var req dto.LoginRequest
	if err := g.ShouldBind(&req); err != nil {
		errorhandler.HandleError(g, &errorhandler.BadRequestError{Message: "Bad Request to server"})
		return
	}

	user, err := h.userService.Login(&req)
	if err != nil {
		errorhandler.HandleError(g, err)
		return
	}

	res := dto.LoginResponse{
		ID:   user.ID,
		Name: user.Name,
	}

	g.JSON(http.StatusOK, res)
}

func (h *authHandler) Register(g *gin.Context) {
	var req dto.RegisterRequest
	if err := g.ShouldBind(&req); err != nil {
		errorhandler.HandleError(g, &errorhandler.BadRequestError{Message: "Invalid request body"})
		return
	}

	if err := h.userService.Register(&req); err != nil {
		errorhandler.HandleError(g, err)
		return
	}

	res:= dto.ResponseParam{
		StatusCode:  http.StatusCreated,
		Message:     "Successfully registered",
	}

	g.JSON(http.StatusCreated, res)
}
