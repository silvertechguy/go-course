package main

import "fmt"

func main() {
	// var name string
	// name = "Ali"
	// age := 23 // Another way to decalre a var
	// fmt.Println("My name is", name, ", and my age is", age)

	// Different value types
	game := "XO"         // string
	players := 2         // int
	gameStarted := false // boolian
	fmt.Println("playing", game, "with", players, "players, game started", gameStarted)

	// default values
	var name string
	var number int
	var floatNumber float32 // float64
	var b bool
	fmt.Println("string default value:", name)
	fmt.Println("int default value:", number)
	fmt.Println("float default value:", floatNumber)
	fmt.Println("boolian default value:", b)

	//  Struct
	type person struct {
		name string
		age  int
	}
	p := person{name: "Ali"}
	fmt.Println(p.name, p.age)

	// Channels
	messages := make(chan string)
	go func() {
		messages <- "Hi :)"
	}() // send message
	msg := <-messages // read message
	fmt.Println(msg)

}
