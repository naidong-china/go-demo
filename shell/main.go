package main

import (
	"bytes"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("/bin/sh", "-c", `echo hello, data-sync`)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	log.Printf("output: %s", stdout.String())

}
