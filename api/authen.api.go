package api

import (
	"main/db"
	"main/model"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SetupAuthenAPI(router *gin.Engine) {
	authenAPI := router.Group("api/v2")
	{
		authenAPI.POST("/login", login)
		authenAPI.POST("/register", register)
	}
}

// Login - login api
func login(c *gin.Context) {
	var user model.User

	if c.ShouldBind(&user) == nil {
		var queryUser model.User
		if err := db.GetDB().First(&queryUser, "username = ?", user.Username).Error; err != nil {
			c.JSON(200, gin.H{"result": "nok", "error": err})
		} else if checkPasswordHash(user.Password, queryUser.Password) == false {
			c.JSON(200, gin.H{"result": "nok", "error": "invalid password"})
		} else {
			c.JSON(200, gin.H{"result": "ok", "data": user})
		}
	} else {
		c.JSON(401, gin.H{"status": "unable to bind data"})
	}
}

func register(c *gin.Context) {
	var user model.User
	if c.ShouldBind(&user) == nil {
		user.Password, _ = hashPassword(user.Password)
		user.CreatedAt = time.Now()
		if err := db.GetDB().Create(&user).Error; err != nil {
			c.JSON(200, gin.H{"result": "nok", "error": err})
		} else {
			c.JSON(200, gin.H{"result": "ok", "data": user})
		}
	} else {
		c.JSON(401, gin.H{"status": "unable to bind data"})
	}

}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func hashPassword(password string) (string, error) {
	byte, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(byte), err
}
