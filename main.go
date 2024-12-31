package main

import (
	"log"
	"yaml-utilities/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("Error: %v\n", err)
	}
}
