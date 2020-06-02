package main

import (
	"fmt"
	"time"
)

//var words = []int64{39, 68, 45, 64, 30, 30, 32, 58, 55, 56, 48, 4d, 43, 39, 31, 52, 75, 78, 47, 61 ,33 ,6b ,6c, 62, 42, 39, 46, 5a, 78, 51, 7a ,63 ,66 ,42, 44 ,61, 58, 74, 6e, 52, 55, 4e, 30, 63, 79, 39, 47, 64 ,7a ,46 ,32, 59 }
var words = []int64{57, 104, 69, 100, 48, 48, 50, 88, 85, 86, 72, 77, 67, 57, 49, 82, 117, 120, 71, 97, 51, 107, 108, 98, 66, 57, 70, 90, 120, 81, 122, 99, 102, 66, 68, 97, 88, 116, 110, 82, 85, 78, 48, 99, 121, 57, 71, 100, 122, 70, 50, 89}

// var word = []int64{173, 127, 150, 60, 137, 163, 64, 61, 144, 137}
// var letters = []int64{101, 156, 131, 67, 150, 154, 156, 107, 137, 102}
// var letter = []int64{60, 165, 124, 137, 155, 64, 164, 110, 77}

func main() {
	fmt.Println("Estou procurando as palavras para falar em inglês ...")
	time.Sleep(5 * time.Second)

	var sentence []string

	for _, v := range words {
		sentence = append(sentence, speak(uint64(v-50), 50))
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
