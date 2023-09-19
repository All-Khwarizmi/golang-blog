package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/All-Khwarizmi/golang-blog/internal/models"
	jsontoken "github.com/All-Khwarizmi/golang-blog/internal/utils/jsonToken"
	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var user = models.User{}

	errJson := c.BindJSON(&user)
	if errJson != nil {
		log.Fatalf("Unable to bind JSON in login. Error: %v / User: %v \n", errJson, user)

	}
	fmt.Printf("User Name: %s  - USER email: %s \n", user.Name, user.Email)

	token, err := jsontoken.GenerateJWT(user.Email, user.Name)
	if err != nil {
		log.Fatalf("Unable to get token to secure protected routes. Error: %v \n", err)

	}

	c.SetCookie("auth", token, 3600, "/", "localhost", false, true)
	fmt.Println("Cookie set, redirecting....")
	c.Redirect(http.StatusFound, "/protected/home-page")

}
