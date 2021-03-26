package main

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/crypto/ssh"
)

type Connection struct {
	*ssh.Client
}

func main() {
	conn, err := Connect("server.example.com:22", "root", "password")
	if err != nil {
		log.Fatal(err)
	}

	output, err := conn.SendCommands("ls /")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The Output is :")
	fmt.Println(string(output))
}

// connect with the server
func Connect(addr, user, password string) (*Connection, error) {
	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.HostKeyCallback(func(hostname string, remote net.Addr, key ssh.PublicKey) error { return nil }),
	}

	conn, err := ssh.Dial("tcp", addr, sshConfig)
	if err != nil {
		return nil, err
	}

	return &Connection{conn}, nil

}

// run cmd in a server
func (conn *Connection) SendCommands(cmd string) ([]byte, error) {
	session, err := conn.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	err = session.RequestPty("xterm", 80, 40, modes)
	if err != nil {
		return []byte{}, err
	}

	output, err := session.Output(cmd)


	if err != nil {
		return output, fmt.Errorf("failed to execute command '%s' on server: %v", cmd, err)
	}

	return output, err
}
