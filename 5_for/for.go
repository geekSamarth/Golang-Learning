package main

import "fmt"

// important fact for 'for' in golang is that in golang for everyt used in looping we have to use the for keyword
func main() {
	// while loop
	i := 1
	for i <= 3 {
		fmt.Println("i: ", i)
		i++
	}

	// infinite loop

	// for {
	// 	fmt.Println("1")
	// }

	// classic for loop

	for i := 0; i <= 3; i++ {
		fmt.Println("i: ", i)
	}

	// break and continue in looping
	for i := 0; i <= 3; i++ {

		// continue block
		// if i == 2 {
		// 	continue
		// }

		// break block
		if i == 2 {
			break
		}
		fmt.Println("i: ", i)
	}

}
