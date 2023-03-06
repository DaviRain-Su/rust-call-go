package main

import "C"
import (
	"fmt"
)

//export HelloWorld
func HelloWorld() *C.char {
	return C.CString("Hello, world, From GO!")
}

func main() {
	fmt.Printf("Hello, world!\n")
}
