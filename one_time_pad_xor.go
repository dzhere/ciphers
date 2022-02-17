package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomBits(length int) []int {
	var result []int = make([]int, length)

	for i := 0; i < length; i++ {
		result = append(result, rand.Intn(2))
	}

	return result
}

func OneTimePad(message, key []int) []int {
	length := len(message)
	if len(message) != len(key) {
		return nil
	}

	var result []int = make([]int, length)
	for i := 0; i < length; i++ {
		if message[i] == key[i] {
			result[i] = 0
			continue
		}

		result[i] = 1
	}

	return result
}

func main() {
	message := RandomBits(25)
	key := RandomBits(25)
	ciphertext := OneTimePad(message, key)
	fmt.Println(reflect.DeepEqual(message, OneTimePad(ciphertext, key)))
}
