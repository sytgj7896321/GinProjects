package main

import (
	"github.com/gin-gonic/gin"
)

func download(c *gin.Context) {
	fileName := c.Query("fileName")
	c.File(fileDir + fileName)
}
