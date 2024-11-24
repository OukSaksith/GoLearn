package main

import "fmt"

func main() {

	//initial map
	pcStock := map[string]int{
		"therog":  2,
		"macbook": 3,
	}

	//add element
	pcStock["lenovo"] = 4

	//update element
	pcStock["therog"] = 7

	//access element
	fmt.Println("Stock the TheROG : ", pcStock["therog"])

	//check if key is exists
	if val, exists := pcStock["hp"]; exists {
		fmt.Println("Stock the HP : ", val)
	} else {
		fmt.Println("HP not found.")
	}

	//delete an element
	delete(pcStock, "hp")

	//Iterator over a map
	fmt.Println("PC Stock after deletion : ")
	for pc, stock := range pcStock {
		fmt.Printf("%s : %d\n", pc, stock)
	}

	//Get the size of map
	fmt.Println("Total of PC type is : ", len(pcStock))

}
