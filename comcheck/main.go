package main

import (
	"fmt"
	"os"
)

// os

func main() {
	for _, arg := range os.Args[1:] {
		if arg == "01" || arg == "galaxy" || arg == "galaxy 01" {
			fmt.Print("Alert!!!\n")
			return
		}
	}
}
