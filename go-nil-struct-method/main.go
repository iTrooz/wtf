package main

import (
	"fmt"
)

type MyStruct struct {
}

func (m *MyStruct) Foo() {
	fmt.Printf("%v\n", m) // prints "<nil>"
}

func main() {
	var ms *MyStruct = nil
	ms.Foo()
}
