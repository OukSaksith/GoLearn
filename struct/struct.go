package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {

	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Fong", age: 30})

	fmt.Println(person{name: "Yi Ling"})

	fmt.Println(&person{name: "San", age: 40})

	fmt.Println(newPerson("Pheap"))

	s := person{name: "Man", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"Co",
		true,
	}
	fmt.Println(dog)
}
