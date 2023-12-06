package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/karanr1990/go-jwt-app/middleware"
	"net/http"
	"time"
)

func main() {
	router := gin.Default()
	router.POST("/login", loginHandler)
	router.GET("/protected", middleware.AuthMiddleware(), protectedHandler)
	router.Run(":8080")
}
func loginHandler(c *gin.Context) {
	// Mocking the authentication process
	username := c.PostForm("username")
	password := c.PostForm("password")
	// Perform your authentication logic here, e.g., validate the credentials against a database
	// If authentication succeeds, generate a JWT token
	if username == "admin" && password == "admin123" {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["user"] = username
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expiration time
		tokenString, _ := token.SignedString([]byte("your_secret_key"))
		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}
func protectedHandler(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{"message": "Hello, " + user.(string)})
}
