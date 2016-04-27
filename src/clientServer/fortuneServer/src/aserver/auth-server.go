/*
Authentication Server

Usage:
$ go run auth-server.go [aserver UDP ip:port] [fserver RPC ip:port] [secret]

Example:
export GOPATH=<PROJECT_DIRECTORY>
cd PROJECT_DIRECTORY
go run src/aserver/auth-server.go localhost:16210 localhost:16806 2016
*/

package main

import (
	"clientServerUtils"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"net/rpc"
	"os"
	"strconv"
)

//Generate new nonce for udpAddr and add to nonceMap
func generateNewNonce(udpAddr *net.UDPAddr, nonceMap clientServerUtils.ConcurrentMap) int64 {
	nonce := rand.Int63()
	nonceMap.Lock()
	nonceMap.Map[udpAddr.String()] = nonce
	nonceMap.Unlock()
	return nonce
}

func sendNewNonce(conn net.UDPConn, sendto *net.UDPAddr, nonceMap clientServerUtils.ConcurrentMap) {
	nonce := generateNewNonce(sendto, nonceMap)

	nonceMsg := clientServerUtils.NonceMessage{nonce}
	req, err := json.Marshal(nonceMsg)
	if err != nil {
		fmt.Println("Error marshalling NonceMessage: ", err)
		os.Exit(-1)
	}
	_, err = conn.WriteToUDP(req, sendto)
	if err != nil {
		fmt.Printf("Error writing NonceMessage to client %v: %v\n", sendto, err)
	}
}

func sendErrMessage(conn net.UDPConn, sendto *net.UDPAddr, errMsg clientServerUtils.ErrMessage) {
	req, err := json.Marshal(errMsg)
	if err != nil {
		fmt.Println("Error marshalling ErrMessage: ", err)
		os.Exit(-1)
	}
	_, err = conn.WriteToUDP(req, sendto)
	if err != nil {
		fmt.Println("Error writing ErrMessage to client: ", err)
	}
}

func retrieveFortuneInfo(clientUDPAddr string, fserverRCPIpPort string) clientServerUtils.FortuneInfoMessage {
	rpcClient, err := rpc.Dial("tcp", fserverRCPIpPort)
	if err != nil {
		fmt.Println("Error dialing fserver:", err)
		os.Exit(-1)
	}

	var fortuneInfoMsg clientServerUtils.FortuneInfoMessage
	rpcErr := rpcClient.Call("FortuneServerRPC.GetFortuneInfo", clientUDPAddr, &fortuneInfoMsg)
	if rpcErr != nil {
		fmt.Println("Error in rpcClient.Call:", rpcErr)
		os.Exit(-1)
	}
	rpcClient.Close()
	return fortuneInfoMsg
}

func sendFortuneInfoMessage(conn net.UDPConn, sendto *net.UDPAddr, fortuneInfoMsg clientServerUtils.FortuneInfoMessage) {
	req, err := json.Marshal(fortuneInfoMsg)
	if err != nil {
		fmt.Println("Error marshalling FortuneInfoMessage: ", err)
		os.Exit(-1)
	}
	_, err = conn.WriteToUDP(req, sendto)
	if err != nil {
		fmt.Println("Error writing ErrMessage to client: ", err)
	}
}

// Get FortuneInfoMessage and send to client
func replyWithFortuneInfoMessage(conn net.UDPConn, clientUDPAddr *net.UDPAddr, fserverRCPIpPort string) {
	fortuneInfoMsg := retrieveFortuneInfo(clientUDPAddr.String(), fserverRCPIpPort)
	sendFortuneInfoMessage(conn, clientUDPAddr, fortuneInfoMsg)
}

func parseSecretArg(secretStr string) int64 {
	secret, err := strconv.ParseInt(secretStr, 10, 64)
	if err != nil {
		fmt.Println("Error parsing int from secret arg: ", secretStr, ", Error: ", err)
		os.Exit(-1)
	}
	return secret
}

func isValidHashMessage(received clientServerUtils.HashMessage, storedNonce int64, secret int64) bool {
	expected := clientServerUtils.ComputeHashMessage(storedNonce, secret)
	return expected == received
}

func handleUDPConn(udpListener net.UDPConn, clientUDPAddr *net.UDPAddr, fserverRCPIpPort string, nonceMap clientServerUtils.ConcurrentMap, secret int64, msgFromClient []byte) {
	nonceMap.RLock()
	clientNonce := nonceMap.Map[clientUDPAddr.String()]
	nonceMap.RUnlock()

	var receivedHashMsg clientServerUtils.HashMessage
	err := json.Unmarshal(msgFromClient, &receivedHashMsg)
	if err != nil {
		sendNewNonce(udpListener, clientUDPAddr, nonceMap)
	} else if clientNonce == 0 {
		errMsg := clientServerUtils.ErrMessage{"unknown remote client address"}
		sendErrMessage(udpListener, clientUDPAddr, errMsg)
	} else if isValidHashMessage(receivedHashMsg, clientNonce, secret) {
		replyWithFortuneInfoMessage(udpListener, clientUDPAddr, fserverRCPIpPort)
	} else {
		errMsg := clientServerUtils.ErrMessage{"unexpected hash value"}
		sendErrMessage(udpListener, clientUDPAddr, errMsg)
	}
}

func main() {
	args := os.Args
	if len(args) != 4 {
		fmt.Println("Usage: [aserver UDP ip:port] [fserver RPC ip:port] [secret]")
		os.Exit(-1)
	}
	aserverUDPIpPort := args[1]
	fserverRCPIpPort := args[2]
	secret := parseSecretArg(args[3])

	udpListener := clientServerUtils.InitUDPConn(aserverUDPIpPort)

	//Initialize hash table of (udpIpPort, nonce) key-value pairs
	var nonceMap clientServerUtils.ConcurrentMap
	nonceMap.Map = make(map[string]int64)

	//Start listen/receive connection loop
	for {
		var buf [1024]byte
		msgLen, clientUDPAddr, err := udpListener.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Println("Error on ReadFromUDP: ", err)
		} else {
			go handleUDPConn(udpListener, clientUDPAddr, fserverRCPIpPort, nonceMap, secret, buf[0:msgLen])
		}
	}
	udpListener.Close()
}
