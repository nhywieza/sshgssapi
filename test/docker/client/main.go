package main

import (
	"log"

	"github.com/nhywieza/sshgssapi"
	"golang.org/x/crypto/ssh"
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
