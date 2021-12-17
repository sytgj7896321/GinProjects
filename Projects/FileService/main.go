package main

import (
	"GinProjects/Database"
	"GinProjects/myLibary"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var (
	fileDir = "./files/"
	db      = &Database.Mysql{}
)

func main() {
	db.NewConnection()
	defer func(DB *sql.DB) {
		err := DB.Close()
		myLibary.FailOnError(err, "Fail to close database connection")
	}(db.DB)
	err := db.DB.Ping()
	myLibary.FailOnError(err, "Database connection testing failed")

	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	os.Mkdir(fileDir, 0775)
	os.Chmod(fileDir, 0775)
	router.POST("/upload", upload)
	router.GET("/download", download)
	router.Run(":8080")
}
