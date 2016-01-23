package common

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"strconv"
)

type ErrMessage struct {
	Error string
}

type NonceMessage struct {
	Nonce int64
}

type HashMessage struct {
	Hash string
}

type GoalMessage struct {
	Goal string
}

func ParseIntFromStr(str string) int64 {
	integer, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Printf("Error parsing integer from %s, Error: %v\n", str, err)
		os.Exit(-1)
	}
	return integer
}

func ComputeHashMessage(nonce int64, secret int64) HashMessage {
	dataInt64 := nonce + secret

	dataBuf := make([]byte, 1024)
	length := binary.PutVarint(dataBuf, dataInt64)
	trimmedBuf := dataBuf[:length]

	hash := md5.New()
	hash.Write(trimmedBuf)
	hashStrHex := hex.EncodeToString(hash.Sum(nil))
	return HashMessage{hashStrHex}
}

func ResolveUDPAddr(ipPort string) net.UDPAddr {
	udpAddr, err := net.ResolveUDPAddr("udp", ipPort)
	if err != nil {
		fmt.Printf("Error resolving UDP address %s: %v", ipPort, err)
		os.Exit(-1)
	}
	return *udpAddr
}

// Initialize listener for incoming UDP messages to ip:port
func InitUDPConn(ipPort string) net.UDPConn {
	udpAddr := ResolveUDPAddr(ipPort)

	listener, err := net.ListenUDP("udp", &udpAddr)
	if err != nil {
		fmt.Println("Error initializing aserver UDP listener: ", err)
		os.Exit(-1)
	}
	return *listener
}
