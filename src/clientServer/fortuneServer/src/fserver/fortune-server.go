/*
Fortune Server

Usage:
$ go run fortune-server.go [fserver RPC ip:port] [fserver UDP ip:port] [fortune-string]

Example:
export GOPATH=<PROJECT_DIRECTORY>
cd PROJECT_DIRECTORY
go run src/fserver/fortune-server.go localhost:16806 localhost:15826 "MyFortune"
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
)

type FortuneServerRPC struct {
}

var state struct {
	localUDPAddr string
	nonceMap     clientServerUtils.ConcurrentMap
	fortune      string
}

func (this *FortuneServerRPC) GetFortuneInfo(clientAddr string, fInfoMsg *clientServerUtils.FortuneInfoMessage) error {
	nonce := rand.Int63()
	state.nonceMap.Lock()
	state.nonceMap.Map[clientAddr] = nonce
	state.nonceMap.Unlock()

	fInfoMsg.FortuneNonce = nonce
	fInfoMsg.FortuneServer = state.localUDPAddr
	return nil
}

func resolveTCPAddr(ipPort string) net.TCPAddr {
	tcpAddr, err := net.ResolveTCPAddr("tcp", ipPort)
	if err != nil {
		fmt.Println("Error on ResolveTCPAddr for (ip:port): ", ipPort, ", Error:", err)
		os.Exit(-1)
	}
	return *tcpAddr
}

func sendFortuneMessage(conn net.UDPConn, sendto *net.UDPAddr, fortune string) {
	fortuneMsg := clientServerUtils.FortuneMessage{fortune}
	req, err := json.Marshal(fortuneMsg)
	if err != nil {
		fmt.Println("Error marshalling FortuneMessage: ", err)
		os.Exit(-1)
	}
	_, err = conn.WriteToUDP(req, sendto)
	if err != nil {
		fmt.Println("Error writing FortuneMessage to client: ", err)
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

func isValidNonce(clientUDPAddr *net.UDPAddr, nonceFromClient int64) bool {
	clientUDPIpPort := clientUDPAddr.String()
	state.nonceMap.RLock()
	isValid := state.nonceMap.Map[clientUDPIpPort] == nonceFromClient
	state.nonceMap.RUnlock()
	return isValid
}

func handleClientUDPConn(listener net.UDPConn, clientUDPAddr *net.UDPAddr, msg []byte) {
	//Parse FortuneReqMessage from client
	var fortuneReq clientServerUtils.FortuneReqMessage
	err := json.Unmarshal(msg, &fortuneReq)
	if err != nil {
		errMessage := clientServerUtils.ErrMessage{"could not interpret message"}
		sendErrMessage(listener, clientUDPAddr, errMessage)
	}
	nonceFromClient := fortuneReq.FortuneNonce
	isValid := isValidNonce(clientUDPAddr, nonceFromClient)
	if isValid {
		sendFortuneMessage(listener, clientUDPAddr, state.fortune)
	} else {
		errMessage := clientServerUtils.ErrMessage{"incorrect fortune nonce"}
		sendErrMessage(listener, clientUDPAddr, errMessage)
	}
}

func startRCPRoutine(localRCPIpPort string) {
	fserverRPC := new(FortuneServerRPC)
	rpc.Register(fserverRPC)

	tcpAddr := resolveTCPAddr(localRCPIpPort)
	tcpListener, err := net.ListenTCP("tcp", &tcpAddr)
	if err != nil {
		fmt.Println("Error initializing fserver TCP listener: ", err)
		os.Exit(-1)
	}
	for {
		tcpConnPtr, err := tcpListener.AcceptTCP()
		if err != nil {
			fmt.Println("Error on AcceptTCP: ", err)
			os.Exit(-1)
		}
		rpc.ServeConn(tcpConnPtr)
	}
	tcpListener.Close()
}

//go run fortune-server.go [fserver RPC ip:port] [fserver UDP ip:port] [fortune-string]
func main() {
	args := os.Args
	if len(args) != 4 {
		fmt.Println("Usage: fortune-server [fserver RPC ip:port] [fserver UDP ip:port] [fortune-string]")
		os.Exit(-1)
	}

	//Initialize hash table of (udpIpPort, nonce) key-value pairs
	state.localUDPAddr = args[2]
	state.nonceMap.Map = make(map[string]int64)
	state.fortune = args[3]

	//Initialize TCP routine
	go startRCPRoutine(args[1])

	//Start UDP listen/receive connection loop
	fserverUDPAddr := clientServerUtils.ResolveUDPAddr(state.localUDPAddr)
	udpListener, err := net.ListenUDP("udp", &fserverUDPAddr)
	if err != nil {
		fmt.Println("Error initializing fserver UDP listener: ", err)
		os.Exit(-1)
	}
	for {
		var buf [1024]byte
		msgLen, clientUdpAddr, err := udpListener.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Println("Error on ReadFromUDP: ", err)
		} else {
			go handleClientUDPConn(*udpListener, clientUdpAddr, buf[0:msgLen])
		}
	}
	udpListener.Close()
}
