package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		log.Println(c.Errors.String())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": c.Errors.String(),
		})
	}
}
