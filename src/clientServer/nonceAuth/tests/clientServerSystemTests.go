package main

import (
	aserver "clientServer/nonceAuth/authServer"
	client "clientServer/nonceAuth/client"
	"clientServer/nonceAuth/common"
	"fmt"
	"time"
)

func startAserver(aserverIpPort string, secret int64) {
	go aserver.RunAuthServer(aserverIpPort, secret)

	//Give time for aserver to set up
	time.Sleep(time.Millisecond * 1500)
}

func getNonceFromAserver(clientIpPort string, aserverIpPort string) int64 {
	clientConn := common.InitUDPConn(clientIpPort)
	aserverUDPAddr := common.ResolveUDPAddr(aserverIpPort)
	nonce := client.RetrieveNonce(clientConn, &aserverUDPAddr)
	clientConn.Close()
	return nonce.Nonce
}

func testAssignsNewNonceForNewAddr() {
	firstClientAddr := "localhost:56146"
	secondClientAddr := "localhost:56147"
	aserverAddr := "localhost:56149"
	var secret int64 = 123456

	startAserver(aserverAddr, secret)

	firstNonce := getNonceFromAserver(firstClientAddr, aserverAddr)
	secondNonce := getNonceFromAserver(secondClientAddr, aserverAddr)
	if firstNonce == secondNonce {
		fmt.Println("FAIL, aserver should assign new nonces for connections from different addresses")
	} else {
		fmt.Println("PASS, aserver assigns new nonce for connections from different addresses")
	}
}

func testAssignsNewNonceForSameAddr() {
	clientAddr := "localhost:56146"
	aserverAddr := "localhost:56147"
	var secret int64 = 123456

	startAserver(aserverAddr, secret)

	nonceFromFirstConn := getNonceFromAserver(clientAddr, aserverAddr)
	nonceFromSubsequentConn := getNonceFromAserver(clientAddr, aserverAddr)
	if nonceFromFirstConn == nonceFromSubsequentConn {
		fmt.Println("FAIL, aserver should assign new nonce when make new connection from same address")
	} else {
		fmt.Println("PASS, aserver assigns new nonce for new connection from same address")
	}
}

func main() {
	testAssignsNewNonceForNewAddr()
	testAssignsNewNonceForSameAddr()
}
