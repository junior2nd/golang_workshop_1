package api

import (
	// "hash"

	"github.com/gin-gonic/gin"
	// "golang.org/x/crypto/bcrypt"
)

func SetupAuthenAPI(router *gin.Engine) {
	authenAPI := router.Group("api/v2")
	{
		authenAPI.POST("/login",login)
		authenAPI.POST("/register",register)
	}
}

// Login - login api
func login(c *gin.Context){
	c.JSON(401, gin.H{"status":"login"})
}

func register(c *gin.Context){
	c.JSON(401, gin.H{"status":"register"})
}

// func CheckPasswordHash(password, has string) bool{
// 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// 	return err == nil 
// }