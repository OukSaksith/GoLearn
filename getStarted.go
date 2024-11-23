package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Go" + " Lang")
	fmt.Println(" 1 + 1 = ", 1+1)
	fmt.Println(" 7.2/3.2 = ", 7.2/3.2)
	fmt.Println(" true && false =", true && false)
	fmt.Println(" true || false =", true || false)
	fmt.Println(" !true =", !true)

	var xS = "ValueString"
	fmt.Println("String datatype = ", xS)

	var xAge int = 20
	fmt.Println("Define xAge = ", xAge)

	var xA, xB int = 13, 45
	fmt.Println("Define xA = ", xA, "and define xB = ", xB)

	xC := "GO Lang"
	fmt.Println("No use var xC= ", xC)

	var xE int
	fmt.Println("No assign value = ", xE)

	var linkExample = "https://gobyexample.com/"
	fmt.Println("link = ", linkExample)

	// TO RUN GO : go run file.go
	// TO BUILD GO : go build file.go than .\file

}
