package main

import (
	"log"
)

func main() {
	server := NewAPIServer(":8080", nil)
	if errors := server.Run(); errors != nil {
		log.Fatal(errors)
	}
}
