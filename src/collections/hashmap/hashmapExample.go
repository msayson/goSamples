package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//Declare map of string keys to int64 values
	var nonceMap map[string]int64

	//Allocate memory and initialize nonceMap
	nonceMap = make(map[string]int64)

	//Set values
	key1 := "192.0.0.0:1234"
	nonceMap[key1] = rand.Int63()
	fmt.Printf("key-val pair 1: (\"%s\", %d)\n", key1, nonceMap[key1])

	key2 := "192.0.0.0:5000"
	nonceMap[key2] = rand.Int63()
	fmt.Printf("key-val pair 2: (\"%s\", %d)\n", key2, nonceMap[key2])

	//Overwrite value for key1
	newVal1 := rand.Int63()
	nonceMap[key1] = newVal1
	fmt.Printf("Overwriting nonce[\"%s\"] with new value %d\n", key1, newVal1)
	fmt.Printf("After overwrite, nonceMap[\"%s\"]=%d\n", key1, nonceMap[key1])

	//If key not found, map[key] returns value type's zero value
	fmt.Printf("nonceMap[\"UnknownValue\"]=%d\n", nonceMap["UnknownValue"])
}
