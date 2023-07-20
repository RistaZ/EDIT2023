package main

import (
	"hello-world/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	api.InitializeHandlers(router)

	router.Run(":8080")
}
