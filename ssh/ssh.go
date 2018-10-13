package toolsSsh

import (
	"github.com/dwburke/go-tools"
	"github.com/sfreiberg/simplessh"
	"os/user"
)

type Ssh struct {
	Client *simplessh.Client
}

// good for one-and-done command on a server (or if you have to iterate a lot of servers)
func Run(username string, keyfile string, server string, cmd string) (string, error) {
	var client *simplessh.Client
	var err error

	if len(keyfile) == 0 {
		keyfile = tools.HomeDir() + "/.ssh/id_rsa"
	}

	if len(username) == 0 {
		usr, err := user.Current()
		if err != nil {
			return "", err
		}

		username = usr.Username
	}

	if client, err = simplessh.ConnectWithKeyFile(server, "addict", keyfile); err != nil {
		return "", err
	}
	defer client.Close()

	// Now run the commands on the remote machine:
	x, err := client.Exec(cmd)

	if err != nil {
		return "", err
	}

	return string(x), nil
}

// for creating a reusable client for multiple commands with Run() below
func New(username string, server string, keyfile string) (*Ssh, error) {

	if len(keyfile) == 0 {
		keyfile = tools.HomeDir() + "/.ssh/id_rsa"
	}

	if len(username) == 0 {
		usr, err := user.Current()
		if err != nil {
			return nil, err
		}

		username = usr.Username
	}

	ssh := Ssh{}

	var client *simplessh.Client
	var err error

	if client, err = simplessh.ConnectWithKeyFile(server, username, keyfile); err != nil {
		return nil, err
	}

	ssh.Client = client

	return &ssh, nil
}

func (ssh *Ssh) Run(cmd string) (string, error) {

	// Now run the commands on the remote machine:
	x, err := ssh.Client.Exec(cmd)

	if err != nil {
		return "", err
	}

	return string(x), nil
}

func (ssh *Ssh) Close() {
	ssh.Client.Close()
}
