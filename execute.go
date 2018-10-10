package sshUtils

import (
	"fmt"
	"bytes"
	"golang.org/x/crypto/ssh"
)

func ExecuteSSH(cmd string, client *ssh.Client) string {
	session, err := client.NewSession()
	if err != nil {
		panic(err.Error())
	}
	defer session.Close()

	fmt.Println("Executing command\n" + cmd)

	var buff bytes.Buffer
	session.Stdout = &buff
	if err := session.Run(cmd); err != nil {
		panic(err.Error())
	}
	return buff.String()
}
