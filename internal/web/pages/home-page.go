package pages

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "home-page.html", gin.H{
		"title": "Home Page",
	})
}

func ProtectedHome(c *gin.Context) {
	c.HTML(http.StatusOK, "home-page.html", gin.H{
		"title": "Protected Home Page",
	})
}
