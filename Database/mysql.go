package Database

import (
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
)

func init() {
	flag.StringVar(&Username, "username", "root", "")
	flag.StringVar(&Password, "password", "", "")
	flag.StringVar(&EndPoint, "endpoint", "127.0.0.1:3306", "")
	flag.StringVar(&Database, "database", "", "")
	flag.StringVar(&Options, "options", "", "")
	flag.Parse()
	DBConnection = Username + ":" + Password + "@tcp(" + EndPoint + ")/" + Database + "?" + Options
}

func (m *Mysql) NewConnection() {
	db, err := sql.Open("mysql", DBConnection)
	m.DB = db
	commonFunc.FailOnError(err, "Fail to open database connection")
	m.DB.SetMaxOpenConns(100)
	m.DB.SetMaxIdleConns(100)
}
