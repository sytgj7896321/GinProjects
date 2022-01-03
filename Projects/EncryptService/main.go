package main

import (
	"GinProjects/Projects/EncryptService/x25519"
	"GinProjects/database"
	"GinProjects/myLibary"
	"bytes"
	"flag"
	"github.com/dimiro1/banner"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strconv"
)

var (
	servicePort int
)

func init() {
	flag.IntVar(&servicePort, "port", 8080, "-port portNumber")
	database.InitFlag()
	flag.Parse()
}

func main() {
	banner.Init(os.Stdout, true, true, bytes.NewBufferString(myLibary.Banner("Banner2")))

	closeDB := connectDB()
	defer func() {
		err := closeDB()
		myLibary.FailOnError(err, "Fail to close database connection")
	}()

	var s x25519.Secure
	startServer(s)
}

func connectDB() func() error {
	db, err := database.NewDatabaseConnection()
	myLibary.FailOnError(err, "Fail to open database connection")
	err = db.Ping()
	myLibary.FailOnError(err, "Database connection testing failed")
	return db.Close
}

func startServer(s x25519.Secure) {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	router.GET("/getPubKey", s.GenerateKeyPair)

	err := router.Run(":" + strconv.Itoa(servicePort))
	if err != nil {
		log.Fatalf("Start Service Failed, %s\n", err)
	}
}
