package main

import "fmt"

func main() {
	name := getName()
	age := getAge()
	fmt.Println("Your name is", name)
	fmt.Println("You are", age, "years old")
}

func getName() string {
	return "John"
}

func getAge() int {
	return 30
}
