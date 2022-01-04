package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

var (
	privateKeyStr string
	publicKeyStr  string
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			// header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			// 允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")                                                                                                                                                              // 设置返回格式是json
		}
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		c.Next() // 处理请求
	}
}

func StartServer() {
	r := gin.Default()
	r.Use(Cors())
	r.GET("/getPubKey", func(c *gin.Context) {
		if err := GenRsaKey(1024); err != nil {
			log.Println("密钥文件生成失败！")
			c.JSON(500, gin.H{
				"message": "server error",
			})
			return
		}
		c.JSON(200, gin.H{
			"public_key":  publicKeyStr,
			"private_key": privateKeyStr,
		})
	})
	r.POST("/login", func(context *gin.Context) {
		type Param struct {
			EncryptPwd string `json:"encryptPwd"`
		}
		var par Param
		err := context.BindJSON(&par)
		if err != nil {
			log.Println(err.Error())
			context.JSON(200, gin.H{
				"message": "password error",
			})
		}

		//log.Println("par.EncryptPwd is", par.EncryptPwd)
		data, err := base64.StdEncoding.DecodeString(par.EncryptPwd)
		if err != nil {
			log.Println(err.Error())
			return
		}

		passWord, err := RsaDecrypt(data)
		if err != nil {
			log.Println(err.Error())
			return
		}

		//log.Println("password is", string(passWord))
		context.JSON(200, gin.H{
			"message": string(passWord),
		})

	})
	r.POST("/encrypt", func(context *gin.Context) {
		type Pass struct {
			Password string `json:"password"`
		}
		var p Pass
		err := context.BindJSON(&p)
		if err != nil {
			log.Println(err.Error())
			return
		}
		encrypt, err := RsaEncrypt([]byte(p.Password))
		if err != nil {
			log.Println(err.Error())
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"encrypted password": encrypt,
		})
	})
	err := r.Run(":8081") // listen and serve on 0.0.0.0:8081
	if err != nil {
		panic(err)
	}

}

func RsaEncrypt(origData []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode([]byte(publicKeyStr))
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	//解密
	block, _ := pem.Decode([]byte(privateKeyStr))
	if block == nil {
		return nil, errors.New("private key error")
	}
	//解析PKCS1格式的私钥
	private, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	data, err := rsa.DecryptPKCS1v15(rand.Reader, private, ciphertext)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return data, nil
}

func GenRsaKey(bits int) error {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)

	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	priBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}

	privateKeyStr = string(pem.EncodeToMemory(priBlock))
	//fmt.Printf("%v\n", privateKeyStr)
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	publicBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}

	publicKeyStr = string(pem.EncodeToMemory(publicBlock))
	//fmt.Printf("%v\n", publicKeyStr)

	if err != nil {
		return err
	}
	return nil
}
