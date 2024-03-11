package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/revandpratama/go-simple-blog-api/config"
	"github.com/revandpratama/go-simple-blog-api/routes"
)

func main() {
	fmt.Println("Starting config...")
	config.LoadConfig()

	fmt.Println("Database initiated..")
	config.LoadDB()

	r := gin.Default()

	api := r.Group("/api")

	routes.PostRoutes(api)
	routes.UserRouter(api)


	addresStr := fmt.Sprintf("localhost:%v", config.ENV.PORT)

	r.Run(addresStr)
}
