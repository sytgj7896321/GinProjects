package main

import (
	"GinProjects/myLibary"
	"bufio"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func upload(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]
	for _, file := range files {
		_ = c.SaveUploadedFile(file, fileDir+file.Filename)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": fmt.Sprintf("%d files uploaded", len(files)),
	})

}

func uploadImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if myLibary.ImageError(c, err) {
		return
	}
	body, err := file.Open()
	if myLibary.ImageError(c, err) {
		return
	}
	srcImg, err := imaging.Decode(body)
	if myLibary.ImageError(c, err) {
		return
	}
	dstImg := imaging.Resize(srcImg, 55, 55, imaging.Lanczos)
	emptyFile, err := os.OpenFile(fileDir+file.Filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if myLibary.ImageError(c, err) {
		return
	}
	w := bufio.NewWriter(emptyFile)
	defer func(w *bufio.Writer) {
		err := w.Flush()
		if myLibary.ImageError(c, err) {
			return
		}
	}(w)
	err = imaging.Encode(w, dstImg, imaging.JPEG)
	if myLibary.ImageError(c, err) {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "image uploaded",
	})
}
