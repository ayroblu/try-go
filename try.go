package main

import "fmt"

func main() {
	fmt.Printf("%+v\n", 0x2p10)
	fmt.Printf("%+v\n", 0x2.p10)
	fmt.Printf("%+v\n", 0x2p32)
	fmt.Printf("%+v\n", 0x2p64)
	fmt.Printf("%+v\n", 0x2p65)
	fmt.Printf("%+v\n", 0x2p128)
	fmt.Println("Hello, World!")
	t := struct{ Name string }{Name: "Alice"}
	fmt.Println(t) // Output: {Alice}
}
