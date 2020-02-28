package test

import "fmt"

type Person struct {
	Name string
}

func PrintPerson(p *Person) {
	fmt.Println(p.Name)
}
