package main

import (
	"github.com/gin-gonic/gin"
)

func download(c *gin.Context) {
	filename := c.Query("file")
	c.File(fileDir + filename)
}
