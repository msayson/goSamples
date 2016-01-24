/*
Client

Usage:
1. Import the client package
2. Call client.RunClient(clientIpPort string, aserverIpPort string, secret int64)
*/

package client

import (
	"clientServer/nonceAuth/common"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"
)

func RetrieveNonce(clientConn net.UDPConn, aserverUDPAddr *net.UDPAddr) common.NonceMessage {
	//Send arbitrary UDP message to aserver to get nonce
	nonceReq := []byte("Hello aserver!  I'd like a nonce!")
	_, err := clientConn.WriteToUDP(nonceReq, aserverUDPAddr)
	if err != nil {
		fmt.Println("Error writing to aserver: ", err)
		os.Exit(-1)
	}

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
func RetrieveGoalMsg(clientConn net.UDPConn, aserverUDPAddr *net.UDPAddr, hashMsg common.HashMessage) common.GoalMessage {
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

func RunClient(clientIpPort string, aserverIpPort string, secret int64) {
	clientConn := common.InitUDPConn(clientIpPort)
	aserverUDPAddr := common.ResolveUDPAddr(aserverIpPort)

	//Set 10 second timeout for ReadFromUDP
	clientConn.SetReadDeadline(time.Now().Add(10 * time.Second))

	//Retrieve nonce from aserver and compute MD5(nonce + secret)
	fmt.Println("Retrieving nonce from aserver")
	nonceMsg := RetrieveNonce(clientConn, &aserverUDPAddr)
	fmt.Println("Received reply from aserver:", nonceMsg)

	hashMsg := common.ComputeHashMessage(nonceMsg.Nonce, secret)

	//Retrieve GoalMessage from aserver
	fmt.Println("Sending hash to aserver")
	goalMsg := RetrieveGoalMsg(clientConn, &aserverUDPAddr, hashMsg)
	fmt.Println("Received goal message from aserver: ", goalMsg)
}
