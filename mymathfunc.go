// ZER CHANGE

// write a function that...

// adds two numbers
// multiplies two numbers
// divides two numbers
// subtracts

// incorporate error handling

// find a way to include a pointer

package main
import (
	"fmt"
	"strconv"
)

func mySumFunc(x, y int) int {
	return x + y
}

func mySubtractFunc(x, y int) int {
	return x - y
}

func myMultiFunc(x, y int) int {
	return x * y
}

func myDiviFunc(x, y int) int {
	return x / y
}


func getInput(n *int) error {
	var input string
	fmt.Scanln(&input)

	num, err := strconv.Atoi(input)
	if err != nil {
		return fmt.Errorf("input must be an integer")
	}

	*n = num
	return nil
}

// func myErrorMessage() {
// 	var x, y int
// 	//var ValidInput bool

// 	fmt.Println("Please enter an integer, x: ")
// 	if err := getInput(&x); err != nil {
// 		fmt.Println("Error, expecting integer.  Please try again!", err)
// 		return
// 	}

// 	fmt.Println("Please enter an integer, y: ")
// 	if err := getInput(&y); err != nil {
// 		fmt.Println("Error, expecting integer.  Please try again!", err)
// 		return
// 	}

// }

// func myFunkyFuncs(x, y int) int {
// 	mySumFunc := mySumFunc(x, y)
// 	fmt.Printf("The sum of x and y is %d.\n", mySumFunc)

	// mySubFunc := mySubtractFunc(x, y)
	// fmt.Printf("The difference of x and y is %d.\n", mySubFunc)

	// myMultiFunc := myMultiFunc(x, y)
	// fmt.Printf("The multiplication of x and y is %d.\n", myMultiFunc)

	// myDiviFunc := myDiviFunc(x, y)
	// fmt.Printf("The division of x and y is %d.\n", myDiviFunc)

func main() {

	fmt.Println("---------Starting the function!------------")

	var x, y int
	//var ValidInput bool

	fmt.Println("Please enter an integer, x: ")
	if err := getInput(&x); err != nil {
		fmt.Println("Error, expecting integer.  Please try again!", err)
		return
	}

	fmt.Println("Please enter an integer, y: ")
	if err := getInput(&y); err != nil {
		fmt.Println("Error, expecting integer.  Please try again!", err)
		return
	}

	mySumFunc := mySumFunc(x, y)
	fmt.Printf("The sum of x and y is %d.\n", mySumFunc)

	mySubFunc := mySubtractFunc(x, y)
	fmt.Printf("The difference of x and y is %d.\n", mySubFunc)

	myMultiFunc := myMultiFunc(x, y)
	fmt.Printf("The multiplication of x and y is %d.\n", myMultiFunc)

	myDiviFunc := myDiviFunc(x, y)
	fmt.Printf("The division of x and y is %d.\n", myDiviFunc)

	fmt.Println("------------------End Print-----------------")
}

// my change right here