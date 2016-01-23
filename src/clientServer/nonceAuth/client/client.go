/*
Client

Usage:
$ go run client.go [local UDP ip:port] [aserver UDP ip:port] [secret]

Local PC:
go install clientServer/nonceAuth/client
"bin/client" localhost:16543 localhost:15432 123456
*/

package main

import (
	"clientServer/nonceAuth/common"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func retrieveNonce(clientConn net.UDPConn, aserverUDPAddr *net.UDPAddr) common.NonceMessage {
	//Send arbitrary UDP message to aserver to get nonce
	nonceReq := []byte("Hello aserver!  I'd like a nonce!")
	_, err := clientConn.WriteToUDP(nonceReq, aserverUDPAddr)

	//Receive NonceMessage reply
	var buf [1024]byte
	msgLen, err := clientConn.Read(buf[:])
	if err != nil {
		fmt.Println("Error on read: ", err)
		os.Exit(-1)
	}
	nonceReplyStr := string(buf[0:msgLen])
	bufBytes := []byte(nonceReplyStr)

	var nonce common.NonceMessage
	json.Unmarshal(bufBytes, &nonce)
	return nonce
}

//Retrieves GoalMessage from aserver
func retrieveGoalMsg(clientConn net.UDPConn, aserverUDPAddr *net.UDPAddr, hashMsg common.HashMessage) common.GoalMessage {
	//Send HashMessage to aserver to get GoalMessage
	req, err := json.Marshal(hashMsg)
	if err != nil {
		fmt.Println("Error marshalling hashMsg: ", err)
		os.Exit(-1)
	}
	_, err = clientConn.WriteToUDP(req, aserverUDPAddr)
	if err != nil {
		fmt.Println("Error writing hashMsg to aserver: ", err)
		os.Exit(-1)
	}

	//Receive GoalMessage reply
	var buf [1024]byte
	msgLen, err := clientConn.Read(buf[:])
	if err != nil {
		fmt.Println("Error on read: ", err)
		os.Exit(-1)
	}
	var goalMsg common.GoalMessage
	json.Unmarshal(buf[0:msgLen], &goalMsg)
	return goalMsg
}

func main() {
	args := os.Args //os.Args[0] is program name
	if len(args) != 4 {
		fmt.Println("Usage: client [local UDP ip:port] [aserver UDP ip:port] [secret]")
		os.Exit(-1)
	}
	clientIpPort := args[1]
	aserverIpPort := args[2]
	secret := common.ParseIntFromStr(args[3])

	clientConn := common.InitUDPConn(clientIpPort)
	aserverUDPAddr := common.ResolveUDPAddr(aserverIpPort)

	//Retrieve nonce from aserver and compute MD5(nonce + secret)
	fmt.Println("Retrieving nonce from aserver")
	nonceMsg := retrieveNonce(clientConn, &aserverUDPAddr)
	fmt.Println("Received reply from aserver:", nonceMsg)

	hashMsg := common.ComputeHashMessage(nonceMsg.Nonce, secret)

	//Retrieve GoalMessage from aserver
	fmt.Println("Sending hash to aserver")
	goalMsg := retrieveGoalMsg(clientConn, &aserverUDPAddr, hashMsg)
	fmt.Println("Received goal message from aserver: ", goalMsg)
}
