package main

import (

	"fmt"

	"example.com/greetings"
)
	


func main(){
	message := greetings.Hello("metocs")

	fmt.Println(message)
}