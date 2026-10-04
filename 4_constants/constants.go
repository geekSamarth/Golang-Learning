package main

import "fmt"

const name string = "John Doe"
const age int = 20
const isMale bool = true
const salary float64 = 10000.00

// salary:=10000 (not allowed as the salary is currently in global scope)

// constant group
const (
	port = 5000
	host = "localhost"
)

func main() {
	const name string = "John Doe"
	const age int = 20
	const isMale bool = true
	const salary float64 = 10000.00

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isMale)
	fmt.Println(salary)

	// salary =10000  (not allowed as the salary is a constant variable)

	fmt.Println(port, host)
}
