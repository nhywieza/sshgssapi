package main

import (
	"io/ioutil"
	"log"
	"net"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/gss"
)

func main() {
	buf, _ := ioutil.ReadFile("/root/.ssh/id_rsa")
	hostKey, _ := ssh.ParsePrivateKey(buf)
	gssServer, _ := gss.NewSSHGSSAPIServerSide()
	config := &ssh.ServerConfig{
		GSSAPIWithMICConfig: &ssh.GSSAPIWithMICConfig{
			AllowLogin: func(conn ssh.ConnMetadata, srcName string) (*ssh.Permissions, error) {
				log.Printf("User is %s", conn.User())
				log.Printf("krb5 src name is %s", srcName)
				return nil, nil
			},
			Server: gssServer,
		},
	}
	config.AddHostKey(hostKey)
	l, _ := net.Listen("tcp", "0.0.0.0:1234")
	log.Printf("start server \n")
	defer func() {
		if err := recover(); err != nil {
			log.Print(err)
		}
	}()
	for {
		conn, _ := l.Accept()
		_, _, _, err := ssh.NewServerConn(conn, config)
		log.Print(err)
	}
}
