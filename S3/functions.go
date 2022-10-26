package main

import (
	"fmt"
)

func main() {
	greet("Rifki", 20, "Bogor")
}

func greet(name string, age int8, adress string) {
	fmt.Printf("Hello There! My name is %s and I'm %d years old", name, age)
	fmt.Println("I live in", adress)
}
