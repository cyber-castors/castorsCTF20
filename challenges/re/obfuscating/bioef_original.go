package main

import (
	"fmt"
	"time"
)

var words = []int64{143, 141, 163, 164, 157, 162, 163, 103, 124, 106}
var word = []int64{173, 127, 150, 60, 137, 163, 64, 61, 144, 137}
var letters = []int64{101, 156, 131, 67, 150, 154, 156, 107, 137, 102}
var letter = []int64{60, 165, 124, 137, 155, 64, 164, 110, 77}

func main() {
	fmt.Println("Estou procurando as palavras para falar em inglês ...")
	time.Sleep(5 * time.Second)

	var sentence []string

	for _, v := range words {
		sentence = append(sentence, speak(uint64(v-50), 50))
	}
	for _, v := range word {
		sentence = append(sentence, speak(uint64(v-45), 45))
	}
	for _, v := range letters {
		sentence = append(sentence, speak(uint64(v-20), 20))
	}
	for _, v := range letter {
		sentence = append(sentence, speak(uint64(v-35), 35))
	}
	fmt.Println("Aqui vou")
	fmt.Println(sentence)

}

func factorial(i uint64) uint64 {
	if i == 0 {
		return 1
	}
	return factorial(i-1) * i
}

func speak(k uint64, n uint64) string {
	return string(k + n)
}
