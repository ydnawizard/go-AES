package main

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!",name)
	return message
}

func main() {
	var name string = Hello("Andy")
	fmt.Printf("%s\n",name)
}
