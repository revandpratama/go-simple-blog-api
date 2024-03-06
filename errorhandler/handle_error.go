package errorhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/revandpratama/go-simple-blog-api/dto"
	"github.com/revandpratama/go-simple-blog-api/helper"
)

func HandleError(g *gin.Context, err error) {
	var statusCode int

	switch err.(type) {
	case *BadRequestError:
		statusCode = http.StatusBadRequest
	case *UnauthorizedError:
		statusCode = http.StatusUnauthorized
	case *NotFoundError:
		statusCode = http.StatusNotFound
	case *InternalServerError:
		statusCode = http.StatusInternalServerError
	}

	res := helper.Response(dto.ResponseParam{
		StatusCode: statusCode,
		Message:    err.Error(),
	})

	g.JSON(statusCode, res)
}
