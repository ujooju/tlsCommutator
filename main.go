package main

import (
	"fmt"

	"github.com/ujooju/tlsCommutator/server"
)

func main() {
	go server.SetUpServer()
	command := ""
	for fmt.Scan(&command); ; fmt.Scan(&command) {
		if command == "exit" {
			return
		}
	}
}
