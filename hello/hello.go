package main

import (
	"fmt"
	"greetings"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("dev")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	names := []string{"dev1", "dev2", "dev3"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)

}
