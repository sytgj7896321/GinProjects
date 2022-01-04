package myRedis

import (
	"flag"
)

var (
	Address  string
	Password string
)

func InitFlag() {
	flag.StringVar(&Address, "address", "192.168.123.25:6379,192.168.123.26:6379,192.168.123.27:6379,192.168.123.25:6380,192.168.123.26:6380,192.168.123.27:6380", "Usage: -address host1:port,host2:port,host3:port,...hostN:port")
	flag.StringVar(&Password, "password", "", "Usage: -password <password> (default not use)")
}
