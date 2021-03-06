package main

import (
	"GinProjects/KeyValue/myRedis"
	"GinProjects/database"
	"GinProjects/myLibary"
	"bytes"
	"context"
	"database/sql"
	"github.com/dimiro1/banner"
	"github.com/go-redis/redis/v8"
	"github.com/rookie-ninja/rk-boot"
	"github.com/rookie-ninja/rk-boot/gin"
	"os"
	"strings"
)

var (
	DB          *sql.DB
	Pre1        *sql.Stmt
	Pre2        *sql.Stmt
	RedisClient *redis.ClusterClient
)

func init() {
	myRedis.InitFlag()
	database.InitFlag()
}

func main() {
	//Banner
	banner.Init(os.Stdout, true, true, bytes.NewBufferString(myLibary.Banner("Banner2")))

	//Database
	closeDB := connectDB()
	defer func() {
		err := closeDB()
		myLibary.FailOnError(err, "Fail to close database connection")
	}()

	//Redis
	RedisClient = connectRedis()
	defer func(RedisClient *redis.ClusterClient) {
		err := RedisClient.Close()
		if err != nil {
			myLibary.FailOnError(err, "Fail to close redis connection")
		}
	}(RedisClient)

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
	Pre2 = prepare("select `public_key`, `private_key` from keyPair where `public_key` = ?")
	return DB.Close
}

func startServer() {
	boot := rkboot.NewBoot()
	entry := rkbootgin.GetGinEntry("unitrust")
	entry.Router.MaxMultipartMemory = 8 << 20 // 8 MiB
	unitrust := entry.Router.Group("/unitrust")
	unitrust.GET("/x25519PubKey", x25519KeyPair)
	unitrust.POST("/x25519Encrypt", x25519Encrypt)
	unitrust.POST("/x25519Decrypt", x25519Decrypt)

	boot.Bootstrap(context.TODO())
	boot.WaitForShutdownSig(context.TODO())
}

func prepare(sql string) *sql.Stmt {
	pre, err := DB.Prepare(sql)
	if err != nil {
		myLibary.FailOnError(err, "Fail to prepare sql: "+sql)
	}
	return pre
}
