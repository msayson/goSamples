/*
Example runner for authentication server

Usage:
$ go run authServerRunner.go [server UDP ip:port] [secret]
*/

package main

import (
	"clientServer/nonceAuth/authServer"
	"clientServer/nonceAuth/common"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("Usage: authServerRunner [server UDP ip:port] [secret]")
		os.Exit(-1)
	}
	udpIpPort := args[1]
	secret := common.ParseIntFromStr(args[2])
	authServer.RunAuthServer(udpIpPort, secret)
}
