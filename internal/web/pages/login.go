package pages

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Blog's Stret",
	})
}
