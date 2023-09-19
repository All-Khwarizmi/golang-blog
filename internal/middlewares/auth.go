package middlewares

import (
	"log"
	"net/http"
	"time"

	"github.com/All-Khwarizmi/golang-blog/internal/models"
	"github.com/All-Khwarizmi/golang-blog/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()
		cookie, err := ctx.Cookie("auth")
		if err != nil {
			log.Fatal("Unauthorized request")
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized request",
			})
		}
		// Parsing jwt
		token, err := jwt.ParseWithClaims(cookie, &models.UserClaim{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(utils.GoDotEnvVariable("SECRET_KEY")), nil
		})
		if err != nil {
			log.Fatal("Unable to parse token from cookie")
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unable to parse token from cookie",
			})
		}

		if !token.Valid {
			log.Fatal("Invalid token")
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid token",
			})
		}
		log.Printf("Token parsed and valid")

		ctx.Next()

		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := ctx.Writer.Status()
		log.Println(status)
	}
}
