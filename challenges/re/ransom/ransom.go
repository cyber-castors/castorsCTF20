package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func getSeed() int64 {

	resp, err := http.Get("http://192.168.0.2:8081/seed")

	if err != nil {
		fmt.Println(err)
		return 0
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	seed, err := strconv.Atoi(strings.Replace(string(body), "\n", "", -1))

	if err != nil {
		fmt.Println(err)
	}

	return int64(seed)
}

func send(seed int) string {
	body := "{seed:" + strconv.Itoa(seed) + "}"
	resp, err := http.Post("http://192.168.0.2:8081/seed", "application/json", bytes.NewBuffer([]byte(body)))

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	r, _ := ioutil.ReadAll(resp.Body)

	return strings.Replace(string(r), "\"", "", -1)
}

func encrypt() {
	k := rand.Intn(254)

	inFile, err := os.Open("flag.png")

	if err != nil {
		fmt.Println(err)
	}

	defer inFile.Close()

	outFile, err := os.OpenFile("flag.png", os.O_RDWR, 0777)

	if err != nil {
		fmt.Println(err)
	}

	defer outFile.Close()

	buffer := make([]byte, 1024)

	for {
		n, err := inFile.Read(buffer)

		if err != nil && err != io.EOF {
			fmt.Println(err)
		}

		if n == 0 {
			break
		}

		for i := range buffer {
			buffer[i] ^= byte(k)
			k = rand.Intn(254)
		}

		if _, err := outFile.Write(buffer[:n]); err != nil {
			fmt.Println(err)
		}

	}

}

func main() {
	seed := getSeed()
	resp := send(int(seed))

	if seed != 0 && resp == "ok\n" {
		rand.Seed(seed)
	} else {
		rand.Seed(time.Now().UTC().UnixNano())
	}

	encrypt()
}
