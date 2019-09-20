package main

import "fmt"

type Cat struct{}

func (c Cat) Say() string {
	return fmt.Println("Meow")
}

type Dog struct{}

func (d Dog) Say() string {
	fmt.Println("woof")
}

func main() {
	c := Cat{}
	c.Say()
	d := Dog{}
	d.Say()
}
