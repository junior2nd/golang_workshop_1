package main

import "github.com/gin-gonic/gin"

// "main/api"

func main() {
	router := gin.Default()
	router.Static("/images", "./uploaded/images")

	// api.Setup(router)
	router.Run(":8081")

}
