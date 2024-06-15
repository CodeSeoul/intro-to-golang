package main

// Slide 18

import (
	"fmt"
	"reflect"
)

type Creature interface {
	MakeSound() string
}

type Person struct {
	Name string
	Age  uint8
}

type FancyInt int

func (i *FancyInt) PrintMe() {
	fmt.Println("fancy int i:", *i)
}

func (p *Person) MakeSound() string {
	return "Hi, I'm " + p.Name + "!"
}

func main() {
	fmt.Println("structs!")

	myPerson := Person{
		Name: "Friendo",
		Age:  20,
	}
	fmt.Println("Person", myPerson.Name, "aged", myPerson.Age)
	myPersonType := reflect.TypeOf(myPerson).Kind()
	fmt.Println("myPersonType:", myPersonType)

	pointerPerson := &Person{
		Name: "Pointo",
		Age:  30,
	}
	fmt.Println("Person", pointerPerson.Name, "aged", pointerPerson.Age)
	pointerPersonType := reflect.TypeOf(pointerPerson).Kind()
	fmt.Println("pointerPersonType:", pointerPersonType)

	var myFancyInt FancyInt = 10
	myFancyInt.PrintMe()

	var notAPerson Creature = &Person{
		Name: "Totally Not a Person",
		Age:  40,
	}
	fmt.Println("Make sound:", notAPerson.MakeSound())
}
