package main

import (
	"fmt"

	"github.com/revandpratama/go-simple-blog-api/config"
)

func main() {
	fmt.Println("Starting config...")
	config.LoadConfig()
}