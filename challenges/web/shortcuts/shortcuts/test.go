package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("cp", "shortcuts/*", "shortcuts"+"123")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
