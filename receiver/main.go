package main

import (
	"fmt"

	"github.com/id3ntity99/module/entity"
)

func main() {
	person := person.NewPerson("John Doe", "Male", 25)
	x := person.GetOlderUsingPointerRecv(10)
	y := person.GetOlderUsingValueRecv(10)
	fmt.Println(x)
	fmt.Println(y)
}
