package main

import (
	"fmt"
	"github.com/m1ker1n/greetings"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := [...]string{
		"Igor",
		"Ne Igor",
		"Somebody",
	}
	msgs, err := greetings.Hellos(names[:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(msgs)
}
