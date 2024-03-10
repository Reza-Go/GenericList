package main

import (
	"fmt"
	"genericList/generics"
)

type Person struct {
	Name string
	Age  int
}

func main() {

	GenericInt()
	GenericString()
	GenericPerson()

}

func GenericInt() {

	list1 := generics.List[int]{Items: []int{}}
	list1.Add(10)
	list1.Add(20)
	list1.Add(30)

	fmt.Printf("%v\n", list1.Items)

	list1.InsertAt(1, 15)
	fmt.Printf("%v\n", list1.Items)

	list1.RemoveAt(2)
	fmt.Printf("%v\n", list1.Items)

	list1.Remove(30)
	fmt.Printf("%v\n", list1.Items)
}

func GenericString() {
	list1 := generics.List[string]{Items: []string{}}
	list1.Add("Ali")
	list1.Add("Jack")
	list1.Add("Jennifer")

	fmt.Printf("%v\n", list1.Items)

	list1.InsertAt(1, "Hossein")
	fmt.Printf("%v\n", list1.Items)

	list1.RemoveAt(2)
	fmt.Printf("%v\n", list1.Items)

	list1.Remove("Jennifer")
	fmt.Printf("%v\n", list1.Items)
}
func GenericPerson() {
	list1 := generics.List[Person]{Items: []Person{}}
	list1.Add(Person{Name: "Ali", Age: 30})
	list1.Add(Person{Name: "Reza", Age: 20})
	list1.Add(Person{Name: "Sadegh", Age: 10})

	fmt.Printf("%v\n", list1.Items)

	list1.InsertAt(1, Person{Name: "Hadi", Age: 40})
	fmt.Printf("%v\n", list1.Items)

	list1.RemoveAt(2)
	fmt.Printf("%v\n", list1.Items)

	list1.Remove(Person{Name: "Sadegh", Age: 10})
	fmt.Printf("%v\n", list1.Items)
}
