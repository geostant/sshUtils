package sshUtils

import (
	"golang.org/x/crypto/ssh"
	"net"
)

func Session(username, password, hostAddress string) *ssh.Client {
	var config = &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error{
			return nil
		},
	}

	client, err := ssh.Dial("tcp", hostAddress + ":22", config)
	if err != nil {
		panic(err.Error())
	}
	return client
}
