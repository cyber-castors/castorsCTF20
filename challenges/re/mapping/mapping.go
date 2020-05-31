package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

var flag string = "eHpzdG9yc1hXQXtpYl80cjFuMmgxNDY1bl80MXloMF82Ml95MDQ0MHJfNGQxbl9iNXVyMn0="
var customMapping string = "ZYXFGABHOPCDEQRSTUVWIJKLMNabczyxjklmdefghinopqrstuvw5670123489+/_{}"
var normalMapping string = "ABCDEFGHIJKLMNOPQRSTUVWZYXzyxabcdefghijklmnopqrstuvw0123456789+/_{}"

func apply(input string, table map[byte]byte) string {
	result := make([]byte, len(input))

	for i, l := range input {
		result[i] = table[byte(l)]
	}

	encoded := base64.StdEncoding.EncodeToString(result)

	return encoded
}

func check(input string) bool {
	return input == flag
}

func createMappingTable() map[byte]byte {
	table := make(map[byte]byte)

	for i, x := range normalMapping {
		table[byte(x)] = customMapping[i]
	}

	return table
}

func main() {

	table := createMappingTable()

	fmt.Println("Enter Password")
	reader := bufio.NewReader(os.Stdin)

	password, _ := reader.ReadString('\n')
	password = strings.Replace(password, "\n", "", -1)

	password = apply(password, table)
	isCorrect := check(password)

	if isCorrect {
		fmt.Println("Correct!")
	} else {
		fmt.Println("Wrong!")
	}
}
