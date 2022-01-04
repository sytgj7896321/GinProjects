package main

import (
	"GinProjects/KeyValue/myRedis"
	"GinProjects/database"
	"GinProjects/myLibary"
	"bytes"
	"database/sql"
	"flag"
	"github.com/dimiro1/banner"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	DB          *sql.DB
	Pre1        *sql.Stmt
	Pre2        *sql.Stmt
	RedisClient *redis.ClusterClient
	servicePort int
)

func init() {
	flag.IntVar(&servicePort, "port", 8080, "-port portNumber")
	myRedis.InitFlag()
	database.InitFlag()
	flag.Parse()
}

func main() {
	//Banner
	banner.Init(os.Stdout, true, true, bytes.NewBufferString(myLibary.Banner("Banner2")))

	//Redis
	RedisClient = connectRedis()

	//Database
	closeDB := connectDB()
	defer func() {
		err := closeDB()
		myLibary.FailOnError(err, "Fail to close database connection")
	}()

	//API
	startServer()
}

func connectRedis() *redis.ClusterClient {
	addresses := strings.Split(myRedis.Address, ",")
	redisClusterClient := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    addresses,
		Password: myRedis.Password,
	})
	return redisClusterClient
}

func connectDB() func() error {
	var err error
	DB, err = database.NewDatabaseConnection()
	myLibary.FailOnError(err, "Fail to open database connection")
	err = DB.Ping()
	myLibary.FailOnError(err, "Database connection testing failed")
	Pre1 = prepare("insert into keyPair values (?, ?)")
	Pre2 = prepare("select `private_key` from keyPair where `public_key` = ?")
	return DB.Close
}

func startServer() {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	router.GET("/x25519PubKey", x25519KeyPair)
	router.POST("/x25519Encrypt", x25519Encrypt)
	router.POST("/x25519Decrypt", x25519Decrypt)

	err := router.Run(":" + strconv.Itoa(servicePort))
	if err != nil {
		log.Fatalf("Start Service Failed, %s\n", err)
	}
}

func prepare(sql string) *sql.Stmt {
	pre, err := DB.Prepare(sql)
	if err != nil {
		myLibary.FailOnError(err, "Fail to prepare sql: "+sql)
	}
	return pre
}
