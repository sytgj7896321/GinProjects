package main

import (
	"GinProjects/myLibary"
	"bufio"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"github.com/pkg/sftp"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func upload(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]
	for _, file := range files {
		_ = c.SaveUploadedFile(file, fmt.Sprintf("%s/%s", fileDir, file.Filename))
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
	emptyFile, err := os.OpenFile(fmt.Sprintf("%s/%s", fileDir, file.Filename), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
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

func transfer(c *gin.Context) {
	remoteDir := c.Query("remoteDir")
	err := sftpClient.MkdirAll(remoteDir)
	if myLibary.TransferError(c, err) {
		return
	}
	file, err := c.FormFile("file")
	if myLibary.TransferError(c, err) {
		return
	}
	srcFile, err := file.Open()
	defer func(srcFile multipart.File) {
		err := srcFile.Close()
		if myLibary.TransferError(c, err) {
			return
		}
	}(srcFile)
	if myLibary.TransferError(c, err) {
		return
	}
	dstFile, err := sftpClient.Create(fmt.Sprintf("%s/%s", remoteDir, file.Filename))
	defer func(dstFile *sftp.File) {
		err := dstFile.Close()
		if myLibary.TransferError(c, err) {
			return
		}
	}(dstFile)
	if myLibary.TransferError(c, err) {
		return
	}
	readerAll, err := ioutil.ReadAll(srcFile)
	if myLibary.TransferError(c, err) {
		return
	}
	_, err = dstFile.Write(readerAll)
	if myLibary.TransferError(c, err) {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "ok",
	})
}
