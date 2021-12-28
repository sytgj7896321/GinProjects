package mySSH

import (
	"flag"
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"time"
)

var (
	sshUser string
	sshPass string
	sshHost string
	sshPort string
)

func InitSSHFlag() {
	flag.StringVar(&sshUser, "sshUser", "root", "")
	flag.StringVar(&sshPass, "sshPass", "", "")
	flag.StringVar(&sshHost, "sshHost", "127.0.0.1", "")
	flag.StringVar(&sshPort, "sshPort", "22", "")
}

func NewSFTPConnection() (*sftp.Client, error) {
	var (
		addr         string
		clientConfig *ssh.ClientConfig
		sshClient    *ssh.Client
		sftpClient   *sftp.Client
		err          error
	)
	// get auth method

	clientConfig = &ssh.ClientConfig{
		User:            sshUser,
		Auth:            []ssh.AuthMethod{ssh.Password(sshPass)},
		Timeout:         15 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// connect to ssh
	addr = fmt.Sprintf("%s:%s", sshHost, sshPort)

	if sshClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	// create sftp client
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		return nil, err
	}

	return sftpClient, nil
}
