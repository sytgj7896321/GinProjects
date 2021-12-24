package main

import (
	"GinProjects/Database"
	"GinProjects/myLibary"
	"database/sql"
	"flag"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var (
	fileDir = "./files/"
	db      = &Database.Mysql{}
	Port    string
)

func init() {
	flag.StringVar(&Port, "port", "8080", "specify FileService port")
	Database.InitFlag()
	flag.Parse()
}

func main() {
	dbConnect()
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	os.Mkdir(fileDir, 0775)
	os.Chmod(fileDir, 0775)
	router.POST("/upload", upload)
	router.POST("/uploadImage", uploadImage)
	router.GET("/download", download)
	router.GET("/hostname", hostname)
	router.Run(":" + Port)
}

func dbConnect() {
	db.NewConnection()
	defer func(DB *sql.DB) {
		err := DB.Close()
		myLibary.FailOnError(err, "Fail to close database connection")
	}(db.DB)
	err := db.DB.Ping()
	myLibary.FailOnError(err, "Database connection testing failed")
}
