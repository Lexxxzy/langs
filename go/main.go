package main

import "fmt"

// function declaration
func add(a int, b int) int {
	return a + b
}

// function literal
var multiply = func(a int, b int) int {
	return a * b
}

// method declaration
type ExampleStruct struct {
	Value int
}

// method declaration
func (e ExampleStruct) MultiplyBy(factor int) int {
	return e.Value * factor
}

// var_declaration
var x int

// var_spec
var y, z int = 1, 2

// short_var_declaration
a := 42

// const_declaration
const pi = 3.14

// const_spec
const (
	True  = true
	False = false
)

// type_declaration and type_spec
type MyInt int

// assignment_statement
x = 10

// identifier
var someIdentifier = "hello"

// field_identifier
type Person struct {
	Name string
	Age  int
}

// package_identifier
import "math"

// type_identifier
type Alias = ExampleStruct

// type_alias_declaration
type StringAlias = string

// blank_identifier
_, err := fmt.Println("This ignores the first return value")

// label_name
loop:
for i := 0; i < 10; i++ {
	if i > 5 {
		break loop
	}
}

func main() {
	fmt.Println("Addition:", add(3, 5))

	fmt.Println("Multiplication:", multiply(3, 5))

	ex := ExampleStruct{Value: 10}
	fmt.Println("Multiply by factor:", ex.MultiplyBy(3))
}