package main

import (
	"fmt"
	"github.com/m1ker1n/greetings"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	msg, err := greetings.Hello("Test")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(msg)
}
