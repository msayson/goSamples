/*
Sample Client

Usage:
$ go run client.go [local UDP ip:port] [aserver UDP ip:port] [secret]

Example:
export GOPATH=<PROJECT_DIRECTORY>
cd PROJECT_DIRECTORY
go run src/testClients/client.go localhost:16066 localhost:16210 2016
*/

package main

import (
	"clientServerUtils"
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
)

func resolveUDPAddr(ipPort string) net.UDPAddr {
	udpAddr, err := net.ResolveUDPAddr("udp", ipPort)
	if err != nil {
		fmt.Println("Error on ResolveUDPAddr for (ip:port): ", ipPort, ", Error:", err)
		os.Exit(-1)
	}
	return *udpAddr
}

// Initialize listener for incoming UDP messages to ip:port
func initUDPConn(ipPort string) net.UDPConn {
	udpAddr := resolveUDPAddr(ipPort)

	//Initialize listener for incoming UDP messages to client
	listener, err := net.ListenUDP("udp", &udpAddr)
	if err != nil {
		fmt.Println("Error on listen: ", err)
		os.Exit(-1)
	}
	return *listener
}

func retrieveNonce(clientConn net.UDPConn, aserverUDPAddr *net.UDPAddr) clientServerUtils.NonceMessage {
	//Send arbitrary UDP message to aserver to get nonce
	nonceReq := []byte("Hello aserver!  I'd like a nonce!")

	fmt.Println("Sending initial request to aserver") //XXX
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

	fmt.Println("Received reply from aserver:", nonceReplyStr) //XXX

	var nonce clientServerUtils.NonceMessage
	json.Unmarshal(bufBytes, &nonce)
	return nonce
}

func computeHashMessage(nonce int64, secret int64) clientServerUtils.HashMessage {
	dataInt64 := nonce + secret

	dataBuf := make([]byte, 1024)
	length := binary.PutVarint(dataBuf, dataInt64)
	trimmedBuf := dataBuf[:length]

	hash := md5.New()
	hash.Write(trimmedBuf)
	hashStrHex := hex.EncodeToString(hash.Sum(nil))
	return clientServerUtils.HashMessage{hashStrHex}
}

//Retrieves FortuneInfoMessage data required to contact fserver
func retrieveFortuneInfoMsg(clientConn net.UDPConn, aserverUDPAddr *net.UDPAddr, hashMsg clientServerUtils.HashMessage) clientServerUtils.FortuneInfoMessage {
	//Send HashMessage to aserver to get FortuneInfoMessage
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

	//Receive FortuneInfoMessage reply
	var buf [1024]byte
	msgLen, err := clientConn.Read(buf[:])
	if err != nil {
		fmt.Println("Error on read: ", err)
		os.Exit(-1)
	}
	var fortuneInfoMsg clientServerUtils.FortuneInfoMessage
	json.Unmarshal(buf[0:msgLen], &fortuneInfoMsg)
	return fortuneInfoMsg
}

func retrieveFortune(clientConn net.UDPConn, fortuneInfoMsg clientServerUtils.FortuneInfoMessage) clientServerUtils.FortuneMessage {
	//Send FortuneReqMessage to fserver
	fortuneReqMsg := clientServerUtils.FortuneReqMessage{fortuneInfoMsg.FortuneNonce}
	fortuneReq, err := json.Marshal(fortuneReqMsg)
	if err != nil {
		fmt.Println("Error marshalling fortuneReqMsg: ", err)
		os.Exit(-1)
	}

	fmt.Println("Retrieving fortune from fserver")
	fserverUDPAddr := resolveUDPAddr(fortuneInfoMsg.FortuneServer)
	_, err = clientConn.WriteToUDP(fortuneReq, &fserverUDPAddr)
	if err != nil {
		fmt.Println("Error writing to fserver: ", err)
		os.Exit(-1)
	}

	//Receive FortuneMessage reply from fserver
	var buf [1024]byte
	msgLen, err := clientConn.Read(buf[:])
	if err != nil {
		fmt.Println("Error on read: ", err)
		os.Exit(-1)
	}
	fortuneReplyStr := string(buf[0:msgLen])
	fortuneBytes := []byte(fortuneReplyStr)

	var fortune clientServerUtils.FortuneMessage
	json.Unmarshal(fortuneBytes, &fortune)
	return fortune
}

func main() {
	args := os.Args //os.Args[0] is program name
	if len(args) != 4 {
		fmt.Println("Usage: client [local UDP ip:port] [aserver UDP ip:port] [secret]")
		os.Exit(-1)
	}
	clientIpPort := args[1]
	aserverIpPort := args[2]
	secret, err := strconv.ParseInt(args[3], 10, 64)
	if err != nil {
		fmt.Println("Error parsing int from secret arg: ", args[3], ", Error:", err)
		os.Exit(-1)
	}

	clientConn := initUDPConn(clientIpPort)
	aserverUDPAddr := resolveUDPAddr(aserverIpPort)

	//Retrieve nonce from aserver and compute MD5(nonce + secret)
	nonce := retrieveNonce(clientConn, &aserverUDPAddr)

	fmt.Println("Received NonceMessage:", nonce) //XXX

	hashMsg := computeHashMessage(nonce.Nonce, secret)

	//Retrieve FortuneInfoMessage from aserver
	fortuneInfoMsg := retrieveFortuneInfoMsg(clientConn, &aserverUDPAddr, hashMsg)

	//Retrieve FortuneMessage from fserver
	fortune := retrieveFortune(clientConn, fortuneInfoMsg)

	//Print out received fortune string followed by newline
	fmt.Println(fortune.Fortune)
	fmt.Println("")
}
