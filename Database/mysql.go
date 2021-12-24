package Database

import (
	"GinProjects/myLibary"
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
	DB *sql.DB
}

var (
	DBConnection string
	Username     string
	Password     string
	EndPoint     string
	Database     string
	Options      string
	Port         string
)

func InitFlag() {
	flag.StringVar(&Username, "username", "root", "")
	flag.StringVar(&Password, "password", "", "")
	flag.StringVar(&EndPoint, "endpoint", "127.0.0.1:3306", "database ip:port")
	flag.StringVar(&Database, "database", "", "database instance to be used")
	flag.StringVar(&Options, "options", "", "database connection options")
}

func (m *Mysql) NewConnection() {
	DBConnection = Username + ":" + Password + "@tcp(" + EndPoint + ")/" + Database + "?" + Options
	db, err := sql.Open("mysql", DBConnection)
	m.DB = db
	myLibary.FailOnError(err, "Fail to open database connection")
	m.DB.SetMaxOpenConns(100)
	m.DB.SetMaxIdleConns(100)
}
