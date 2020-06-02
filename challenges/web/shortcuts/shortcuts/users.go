package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	os.Setenv("PATH", "/usr/bin/:/sbin")
	//command := `grep "/bin/bash" /etc/passwd`
	cmd := exec.Command("grep", "/bin/bash", "/etc/passwd")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%q", out.String())
}
