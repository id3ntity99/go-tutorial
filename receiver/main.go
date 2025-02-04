package main

import (
	"fmt"

	person "github.com/id3ntity99/module/entity"
)

func main() {
	john := person.NewPerson("John Doe", "Male", 25)

	fmt.Printf("[Info] John is %d years old now\n", john.GetAge())
	ages := john.GetOlderUsingValueRecv(10)
	fmt.Printf("[Info] Using value receiver, John gets old to age of %d years old\n", ages)
	fmt.Printf("[GetAge()] John is %d years old now\n", john.GetAge())

	ages = john.GetOlderUsingPointerRecv(10)
	fmt.Printf("[Info] Using pointer receiver, John gets old to age of %d years old\n", ages)
	fmt.Printf("[GetAge()] John is %d years old now\n", john.GetAge())
}
