package main

import (
	"GinProjects/myLibary"
	"bytes"
	"filippo.io/age"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

type KeyPair struct {
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
}

func x25519KeyPair(c *gin.Context) {
	identity, err := age.GenerateX25519Identity()
	if myLibary.InternalError(c, err) {
		return
	}
	_, err = Pre1.Exec(identity.Recipient().String(), identity.String())
	if myLibary.InternalError(c, err) {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"public_key": identity.Recipient().String(),
	})
}

func x25519Encrypt(c *gin.Context) {
	publicKey, exist := c.GetPostForm("public_key")
	if exist == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "public key not found",
		})
		return
	}
	recipient, err := age.ParseX25519Recipient(publicKey)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "public key parse failed",
		})
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		if err.Error() != "http: no such file" {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "file not found",
			})
			return
		}
	}
	if file.Size > 8<<20 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "file is bigger than 8MB",
		})
		return
	}

	f, err := file.Open()
	if myLibary.InternalError(c, err) {
		return
	}

	src, err := ioutil.ReadAll(f)
	if myLibary.InternalError(c, err) {
		return
	}

	err = f.Close()
	if myLibary.InternalError(c, err) {
		return
	}

	b := new(bytes.Buffer)
	defer b.Reset()
	encrypt, err := age.Encrypt(b, recipient)
	if myLibary.InternalError(c, err) {
		return
	}

	_, err = encrypt.Write(src)
	if myLibary.InternalError(c, err) {
		return
	}

	err = encrypt.Close()
	if myLibary.InternalError(c, err) {
		return
	}

	c.Data(http.StatusOK, "application/octet-stream", b.Bytes())
}

func x25519Decrypt(c *gin.Context) {
	//cmd := RedisClient.Get(context.TODO(), pub.PublicKey)
	//result, err := cmd.Result()
	//if result == "" {
	//	if err.Error() != "redis: nil" {
	//		log.Println(err)
	//	}
	//	row := Pre2.QueryRow(pub.PublicKey)
	//	err := row.Scan(&pub.PrivateKey)
	//	if err != nil {
	//		log.Println(err)
	//		c.JSON(http.StatusInternalServerError, gin.H{
	//			"message": "internal server error",
	//		})
	//		return
	//	}
	//	RedisClient.Set(context.TODO(), pub.PublicKey, pub.PrivateKey, 20*time.Minute)
	//	//c.String(http.StatusOK, pub.PrivateKey)
	//} else {
	//	_, err := c.FormFile("file")
	//	if err != nil {
	//		log.Println(err)
	//		c.JSON(http.StatusBadRequest, gin.H{
	//			"message": "request Content-Type isn't multipart/form-data",
	//		})
	//		return
	//	}
	//	_, err = age.ParseX25519Recipient(result)
	//	if err != nil {
	//		log.Println(err)
	//		c.JSON(http.StatusInternalServerError, gin.H{
	//			"message": "public key parse failed",
	//		})
	//		return
	//	}
	//}
}
