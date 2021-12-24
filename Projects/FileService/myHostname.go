package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func hostname(c *gin.Context) {
	name, _ := os.Hostname()
	c.String(http.StatusOK, name)
}
