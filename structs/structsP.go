package structs

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) save() {
	fmt.Println("Saving", p.Name, "to database")
}

func Struct1() {
	person := Person{
		Name: "John",
		Age:  32,
	}

	fmt.Println(person)
	fmt.Println(person.Name)
	fmt.Println(person.Age)

	person.save()
}