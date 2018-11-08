package main

import (
	"fmt"

	"github.com/prologic/go-linkedlist"
)

func main() {
	xs := &linkedlist.List{}
	fmt.Printf("xs = %s\n", xs.String())

	fmt.Println("appending 3 elements to xs ...")
	xs.Append(1)
	xs.Append(2)
	xs.Append(3)

	fmt.Printf("xs = %s\n", xs.String())

	fmt.Println("reverse the list ...")
	xs.Reverse()
	fmt.Printf("xs = %s\n", xs.String())
}
