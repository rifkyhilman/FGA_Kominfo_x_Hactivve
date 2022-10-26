package main

import "fmt"

func main() {
	greet("Rifki", 20)
}

func greet(name string, age int8) {
	fmt.Printf("Hello There! My name is %s and I'm %d years old", name, age)
}
