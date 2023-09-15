package routes

import (
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

}
func LoginHandler(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Here is the resquest body: %s", jsonData)
	c.Redirect(http.StatusFound, "/home-page")

}
