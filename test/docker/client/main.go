package main

import (
	"log"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/gss"
)

func main() {
	gssClient, _ := gss.NewSSHGSSAPIClientSide()
	config := &ssh.ClientConfig{
		User: "hello",
		Auth: []ssh.AuthMethod{
			ssh.GSSAPIWithMICAuthMethod(
				gssClient, "service",
			),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", "service:1234", config)
	if err != nil {
		log.Fatalf("unable to connect: %v", err)
	}
	defer client.Close()
}
