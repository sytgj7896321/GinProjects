package main

import (
	"GinProjects/myLibary"
	"bytes"
	"context"
	"filippo.io/age"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"time"
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
	exists := RedisClient.Exists(context.TODO(), publicKey).Val()
	if exists != 1 {
		pub := new(KeyPair)
		row := Pre2.QueryRow(publicKey)
		err := row.Scan(&pub.PublicKey, &pub.PrivateKey)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"message": "this public key is not generate by me",
			})
			return
		}
		set := RedisClient.Set(context.TODO(), pub.PublicKey, pub.PrivateKey, 20*time.Minute)
		myLibary.InternalError(c, set.Err())
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
	publicKey, exist := c.GetPostForm("public_key")
	if exist == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "public key not found",
		})
		return
	}
	cmd := RedisClient.Get(context.TODO(), publicKey)
	private, err := cmd.Result()
	if private == "" {
		if err.Error() != "redis: nil" {
			log.Println(err)
		}
		pub := new(KeyPair)
		row := Pre2.QueryRow(publicKey)
		err := row.Scan(&pub.PublicKey, &pub.PrivateKey)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"message": "this public key is not generate by me",
			})
			return
		}
		private := pub.PrivateKey
		if decrypt(c, err, private, publicKey, file, true) {
			return
		}
	} else {
		if decrypt(c, err, private, publicKey, file, false) {
			return
		}
	}
}

func decrypt(c *gin.Context, err error, private string, publicKey string, file *multipart.FileHeader, isSet bool) bool {
	privateKey, err := age.ParseX25519Identity(private)
	if err != nil {
		log.Printf("public key: %s error: %s\n", publicKey, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return true
	}
	if isSet == true {
		set := RedisClient.Set(context.TODO(), publicKey, private, 20*time.Minute)
		myLibary.InternalError(c, set.Err())
	}
	if file.Size > 9<<20 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "file is bigger than 8MB",
		})
		return true
	}
	f, err := file.Open()
	if myLibary.InternalError(c, err) {
		return true
	}
	src := new(bytes.Buffer)
	_, err = io.Copy(src, f)
	if myLibary.InternalError(c, err) {
		return true
	}
	err = f.Close()
	if myLibary.InternalError(c, err) {
		return true
	}
	dst, err := age.Decrypt(src, privateKey)
	if err != nil {
		if err.Error() == "no identity matched any of the recipients" {
			log.Printf("public key: %s error: %s\n", publicKey, err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "public key not match file",
			})
			return true
		} else {
			myLibary.InternalError(c, err)
			return true
		}
	}
	r, err := ioutil.ReadAll(dst)
	if myLibary.InternalError(c, err) {
		return true
	}
	c.Data(http.StatusOK, "application/octet-stream", r)
	return false
}
