package main

import (
	"fmt"
)

type Person struct {
	firstname, lastname string
}

// String satisfies the fmt.Stringer interface for the Person type
func (p Person) String() string {
    return fmt.Sprintf("first name : %s, lastname : %s", p.firstname, p.lastname)
}

func main() {
    p := Person{
        firstname: "John",
        lastname: "Doe",
    }

    fmt.Println(p)
    // output: John Doe <johndoe@example.com>
}