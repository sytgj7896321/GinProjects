package main

import (
	"GinProjects/myLibary"
	"GinProjects/mySSH"
	"bytes"
	"flag"
	"github.com/dimiro1/banner"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/sftp"
	"log"
	"os"
)

var (
	fileDir     string
	servicePort string
	//dbOnOff     string
	sftpClient *sftp.Client
)

func init() {
	//flag.StringVar(&dbOnOff, "dbSwitch", "false", "establish database connection or not")
	flag.StringVar(&fileDir, "fileDir", "./files", "specify a path for save files to local")
	flag.StringVar(&servicePort, "servicePort", "8080", "specify FileService port")
	mySSH.InitSSHFlag()
	//database.InitFlag()
	flag.Parse()
}

func main() {
	banner.Init(os.Stdout, true, true, bytes.NewBufferString(myLibary.Banner("Banner2")))
	//DataBase
	//if dbOnOff == "true" {
	//	db, err := database.NewDatabaseConnection()
	//	myLibary.FailOnError(err, "Fail to open database connection")
	//	defer func(DB *sql.DB) {
	//		err := DB.Close()
	//		myLibary.FailOnError(err, "Fail to close database connection")
	//	}(db)
	//	err = db.Ping()
	//	myLibary.FailOnError(err, "Database connection testing failed")
	//}

	//SFTP
	var err error
	sftpClient, err = mySSH.NewSFTPConnection()
	if err != nil {
		myLibary.FailOnError(err, "Fail to connect to SFTP Server")
	}
	defer func(sftpClient *sftp.Client) {
		err := sftpClient.Close()
		if err != nil {
			myLibary.FailOnError(err, "Fail to close SFTP connection")
		}
	}(sftpClient)

	//Web
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	_, err = os.Stat(fileDir)
	if os.IsNotExist(err) {
		err = os.Mkdir(fileDir, 0775)
		if err != nil {
			log.Println(err)
		}
	}

	err = os.Chmod(fileDir, 0775)
	if err != nil {
		log.Println(err)
	}
	router.POST("/upload", upload)
	router.POST("/uploadImage", uploadImage)
	router.POST("/transfer", transfer)
	router.GET("/download", download)
	router.GET("/hostname", hostname)
	err = router.Run(":" + servicePort)
	if err != nil {
		log.Fatalf("Start Service Failed, %s\n", err)
	}
}
