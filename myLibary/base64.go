package myLibary

import (
	"encoding/base64"
	"io/ioutil"
)

func TransferToBase64(filename string) string {
	file, _ := ioutil.ReadFile(filename)
	dst := base64.StdEncoding.EncodeToString(file)
	return dst
}
