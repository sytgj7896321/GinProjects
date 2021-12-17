package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func upload(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]
	for _, file := range files {
		//body, _ := file.Open()
		//srcImg, _ := imaging.Decode(body)
		//emptyFile, _ := os.OpenFile(fileDir+file.Filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		//w := bufio.NewWriter(emptyFile)
		//_ = imaging.Encode(w, srcImg, imaging.JPEG, imaging.JPEGQuality(25))
		//_ = w.Flush()
		//dstImg := imaging.Resize(srcImg, 55, 55, imaging.Lanczos)
		//imaging.Save(dstImg, fileDir+file.Filename)
		err := c.SaveUploadedFile(file, fileDir+file.Filename)
		log.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": fmt.Sprintf("%d files uploaded", len(files)),
	})

}
