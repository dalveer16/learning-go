package main

import("fmt"
"greetings"
)

func main(){
	message:= greetings.Hello("dev")
	fmt.Println(message)
}