package myLibary

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func ImageError(c *gin.Context, err error) bool {
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "image upload failed",
		})
		return true
	}
	return false
}

func TransferError(c *gin.Context, err error) bool {
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "not ok",
		})
		return true
	}
	return false
}
