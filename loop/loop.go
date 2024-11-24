package main

import "fmt"

func main() {

	// A while-like loop: runs as long as i <= 3
	i := 1

	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for i := range 3 {
		fmt.Println("Rang = ", i)
	}

	// An infinite loop that will break immediately after one iteration
	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
