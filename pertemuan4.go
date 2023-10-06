package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

func sum(numbers ...int) int {
	var total int = 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func newCounter() func() int {
	count := 0

	return func() int {
		count += 1

		return count
	}
}

// struct
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// method => Method is a function that attaches to a type (can be a struct or other data type)
func (e Person) fullName() string {
	return e.FirstName + " " + e.LastName
}

func (p Person) getName() string {
	return p.FirstName + " amazing!"
}

// using pointer no need return value and data type of result
func (p *Person) IncreaseAge() {
	p.Age += 1
}

// /////////////// EXAMPLE ///////////////
type Rect struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Struct Using Pointer Receiver
type Employee struct {
	name   string
	salary int
}

func (e *Employee) changeName(newName string) {
	(*e).name = newName
}

// IMPLEMENT INTERFACE
type calculate interface {
	large() int
	doubleLarge() int
}

type Square struct {
	side int
}

func (s Square) large() int {
	return s.side * s.side
}

func (s Square) doubleLarge() int {
	return s.side * s.side * 2
}

// can be dynamice type data and we can type interface{} or any
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func explain(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("i stored string", strings.ToUpper(i.(string)))
	case int:
		fmt.Println("i stored int", i)
	default:
		fmt.Println("i stored something else", i)
	}
}

// error handling object
func myFunc(i int) (int, error) {
	if i <= 0 {
		return -1, errors.New("should be greater than zero")
	}

	return i, nil
}

func myMethod() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error Message:", err)
		}
	}()

	anOddCondition := true
	if anOddCondition {
		panic("I am panicking")
	}
}

func main() {
	// 4. substring
	value := "cat;dog"
	// take substring from index 4 to length of string.
	substring := value[0:len(value)]
	fmt.Println(substring)

	// 5. Replace
	s := "this[things]I would like to remove"
	t := strings.Replace(s, "[", " ", -1)
	fmt.Printf("%s\n", t)

	// 6. Insert
	p := "green"
	index := 2
	q := p[:index] + "HI" + p[index:]
	fmt.Println(p, q)

	avg := sum(2, 4, 3, 5)
	fmt.Println(avg)

	// Anonymous function
	func() {
		fmt.Println("Welcome! to Alif Page")
	}()

	// variable function
	helloWorld := func() {
		fmt.Println("Welcome! to variable function")
	}

	helloWorld()

	// Anonymous function with argument
	func(sentence string) {
		fmt.Println(sentence)
	}("sentence")

	// closure
	counter := newCounter()
	fmt.Println(counter())
	fmt.Println(counter())

	defer func() {
		fmt.Println("Later")
	}()

	fmt.Println("First")

	// POINTER
	// * operator (dereferencing) =>  Declare pointer variable, Access the value stored in the address
	// & Operator (referencing) => Returns the address of a variable, Access the address of a variable to a pointer
	var name string = "Alif"
	var nameAddress *string = &name

	fmt.Println("name (value) :", name)
	fmt.Println("name (adress) :", &name)
	fmt.Println("nameAddress (value) :", *nameAddress)
	fmt.Println("nameAddress (address) :", nameAddress)

	name = "Syah"

	fmt.Println("name (value) :", name)
	fmt.Println("name (adress) :", &name)
	fmt.Println("nameAddress (value) :", *nameAddress)
	fmt.Println("nameAddress (address) :", nameAddress)

	// Zero Value Pointer <nil>
	number_a := 25
	var number_b *int

	if number_b == nil {
		fmt.Println("number_b is", number_b)
		number_b = &number_a
		fmt.Println("number_b after init : is", *number_b)
	}

	// POINTER WITH BUILT-IN NEW()
	var size = new(int)
	fmt.Printf("Size value is %d \n", *size)
	fmt.Printf("Type is %T \n", size)
	fmt.Printf("Address is %v \n", size)
	*size = 85
	fmt.Println("New size value is", *size)

	// struct implementation
	var Person0 Person
	Person0.FirstName = "Muchson"
	Person0.LastName = "Ibi"
	Person0.Age = 27
	fmt.Println(Person0.FirstName, Person0.LastName, Person0.Age)

	// long declaration with assigned value
	var Person1 = Person{"Aliffio", "Syah", 22}
	fmt.Println(Person1)

	// long declaration with assigned value each name fields
	var Person2 = Person{
		FirstName: "Abidzhar",
		LastName:  "Ramadhan",
		Age:       12,
	}
	fmt.Println(Person2)

	// sort declaration
	Person3 := Person{"Zhio", "Xavier", 7}
	fmt.Println(Person3)

	// short declaration with new keyword
	Person4 := new(Person)
	Person4.FirstName = "Nabila"
	Person4.LastName = "Putri"
	Person4.Age = 19
	fmt.Println(*Person4)
	fmt.Println(Person4.fullName())

	PersonA := Person{"John", "Doe", 25}
	fmt.Printf("%v\n", PersonA)
	fmt.Println(PersonA.getName())

	PersonA.IncreaseAge()
	fmt.Println(PersonA.Age)

	rect := Rect{5.0, 4.0}
	cir := Circle{5.0}
	fmt.Printf("Area of rectangle rect = %0.2f\n", rect.Area())
	fmt.Printf("Area of circle cir = %0.2f\n", cir.Area())

	e := Employee{
		name:   "Ross Geller",
		salary: 1200,
	}

	// e before name change
	fmt.Println("e before name change =", e)
	// create pointer to 'e'
	ep := &e
	// change name
	ep.changeName("Aliffio Syah")
	// e after name change
	fmt.Println("e after name change =", e)

	// interface
	var dimResult calculate
	dimResult = Square{10}
	fmt.Println("large square :", dimResult.large())
	fmt.Println("double large square :", dimResult.doubleLarge())

	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)

	// type assertion
	var secret interface{}

	secret = 2
	var number = secret.(int) * 10
	fmt.Println(secret, "multiple by 10 us :", number)

	secret = []string{"apple", "manggo", "banana"}
	var gruits = strings.Join(secret.([]string), ", ")
	fmt.Println(gruits, "is my favorite fruits")

	explain("Hello World")
	explain(52)
	explain(true)

	// error handling
	result, err := myFunc(-1)
	fmt.Println(result, err)

	myMethod()
}

// // model / object (struct)
// // type User struct {
// // 	id    int
// // 	name  string
// // 	email string
// // }

// // type UserInterface interface {
// // 	changeName() string
// // }

// // func (user User) changeName() string {
// // 	return "Hello " + user.name
// // }

// // func (user *User) changeName() {
// // 	user.name = "Hello " + user.name
// // }

// // func tampil(data interface{}) {
// // 	switch data.(type) {
// // 	case int:
// // 		data = data.(int) + 2
// // 	case string:
// // 		data = data.(string) + " hello"
// // 	}

// // 	fmt.Println(data)
// // }

// func pembagian(a, b int) (int, error) {
// 	// if b == 0 {
// 	// 	return 0, errors.New("b tidak boleh 0")
// 	// }

// 	return a / b, nil
// }

// func main() {
// 	defer func() {
// 		result := recover()
// 		if result != nil {
// 			fmt.Println(result)
// 		}
// 	}()

// 	// defer func() {
// 	// 	fmt.Println("a")
// 	// 	fmt.Println("b")
// 	// }()
// 	// fmt.Println("c")

// 	// defer fmt.Println("a")
// 	// fmt.Println("b")
// 	// fmt.Println("c")

// 	result, err := pembagian(3, 0)

// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(result)
// 	}

// 	// var user model.User

// 	// user.Name = "Me"
// 	// user.Email = "me@me.co.id"

// 	// tampil("alta")
// 	// var age []interface{}
// 	// age = append(age, 1)
// 	// age = append(age, "1")
// 	// age = append(age, "4")
// 	// tampil(2)
// 	// tampil("2")

// 	// var user User
// 	// var userInterface UserInterface

// 	// userInterface = user
// 	// fmt.Println(userInterface.changeName())

// 	// user := User{1, "alif", "alif@gmail.com"}

// 	// user.name = "Alif"
// 	// user = user.changeName()
// 	// user.changeName()
// 	// user.id = 1
// 	// user.email = "alif@gmail.com"

// 	// fmt.Println(user)

// 	// pointer
// 	// var age int = 10
// 	// var ageAddress *int = &age

// 	// fmt.Println(*ageAddress)

// 	// var name string = "alif"
// 	// gantiNama(&name)
// 	// fmt.Println(name)
// }

// // func gantiNama(name *string) {
// // 	*name = "hallo " + *name
// // }
