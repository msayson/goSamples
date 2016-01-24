/*
Example client runner

Usage:
$ go run clientRunner.go [local UDP ip:port] [aserver UDP ip:port] [secret]
*/

package main

import (
	"clientServer/nonceAuth/client"
	"clientServer/nonceAuth/common"
	"fmt"
	"os"
)

func main() {
	args := os.Args //os.Args[0] is program name
	if len(args) != 4 {
		fmt.Println("Usage: clientRunner [local UDP ip:port] [aserver UDP ip:port] [secret]")
		os.Exit(-1)
	}
	clientIpPort := args[1]
	aserverIpPort := args[2]
	secret := common.ParseIntFromStr(args[3])

	client.RunClient(clientIpPort, aserverIpPort, secret)
}
