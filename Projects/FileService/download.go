package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func download(c *gin.Context) {
	filename := c.Query("file")
	c.File(fmt.Sprintf("%s/%s", fileDir, filename))
}
