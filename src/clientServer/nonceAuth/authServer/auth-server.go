/*
Authentication Server

Usage:
1. Import the aserver package
2. Call authServer.RunAuthServer(udpIpPort string, secret int64)
*/

package authServer

import (
	"clientServer/nonceAuth/common"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"os"
	"sync"
)

type concurrentMap struct {
	sync.RWMutex
	m map[string]int64
}

//Generate new nonce for udpAddr and add to nonceMap
func generateNewNonce(udpAddr *net.UDPAddr, nonceMap concurrentMap) int64 {
	nonce := rand.Int63()
	nonceMap.Lock()
	nonceMap.m[udpAddr.String()] = nonce
	nonceMap.Unlock()
	return nonce
}

func sendNewNonce(conn net.UDPConn, clientUDPAddr *net.UDPAddr, nonceMap concurrentMap) {
	nonce := generateNewNonce(clientUDPAddr, nonceMap)

	nonceMsg := common.NonceMessage{nonce}
	req, err := json.Marshal(nonceMsg)
	if err != nil {
		fmt.Println("Error marshalling NonceMessage: ", err)
		os.Exit(-1)
	}
	_, err = conn.WriteToUDP(req, clientUDPAddr)
	if err != nil {
		fmt.Printf("Error writing NonceMessage to client %v: %v\n", clientUDPAddr, err)
	}
}

func sendErrMessage(conn net.UDPConn, clientUDPAddr *net.UDPAddr, errMsg common.ErrMessage) {
	req, err := json.Marshal(errMsg)
	if err != nil {
		fmt.Println("Error marshalling ErrMessage: ", err)
		os.Exit(-1)
	}
	_, err = conn.WriteToUDP(req, clientUDPAddr)
	if err != nil {
		fmt.Println("Error writing ErrMessage to client: ", err)
	}
}

func sendGoalMessage(conn net.UDPConn, clientUDPAddr *net.UDPAddr) {
	goalMsg := common.GoalMessage{"You reached the goal!"}
	req, err := json.Marshal(goalMsg)
	if err != nil {
		fmt.Println("Error marshalling GoalMessage: ", err)
		os.Exit(-1)
	}
	_, err = conn.WriteToUDP(req, clientUDPAddr)
	if err != nil {
		fmt.Println("Error writing GoalMessage to client: ", err)
	}
}

func isValidHashMessage(received common.HashMessage, storedNonce int64, secret int64) bool {
	expected := common.ComputeHashMessage(storedNonce, secret)
	return expected == received
}

func handleUDPConn(conn net.UDPConn, clientUDPAddr *net.UDPAddr, nonceMap concurrentMap, secret int64, msgFromClient []byte) {
	nonceMap.RLock()
	clientNonce := nonceMap.m[clientUDPAddr.String()]
	nonceMap.RUnlock()

	var receivedHashMsg common.HashMessage
	err := json.Unmarshal(msgFromClient, &receivedHashMsg)
	if err != nil {
		//Handle mangled message as request for nonce
		sendNewNonce(conn, clientUDPAddr, nonceMap)
	} else if clientNonce == 0 {
		errMsg := common.ErrMessage{"Unknown client address"}
		sendErrMessage(conn, clientUDPAddr, errMsg)
	} else if isValidHashMessage(receivedHashMsg, clientNonce, secret) {
		sendGoalMessage(conn, clientUDPAddr)
	} else {
		errMsg := common.ErrMessage{"Unexpected hash value"}
		sendErrMessage(conn, clientUDPAddr, errMsg)
	}
}

func RunAuthServer(udpIpPort string, secret int64) {
	udpConn := common.InitUDPConn(udpIpPort)

	//Initialize hash table of (udpIpPort, nonce) key-value pairs
	var nonceMap concurrentMap
	nonceMap.m = make(map[string]int64)

	//Start listen/receive connection loop
	for {
		var buf [1024]byte
		msgLen, clientUDPAddr, err := udpConn.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Println("Error on ReadFromUDP: ", err)
		} else {
			//Start go routine to handle client, continue listening for new clients
			go handleUDPConn(udpConn, clientUDPAddr, nonceMap, secret, buf[0:msgLen])
		}
	}
	udpConn.Close()
}
