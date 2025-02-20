package myutils

import "fmt"

// Global Variable (Public)
var Public string = "Public Variable"

// Global Variable (Private)
var privateVar int = 42

// Constant
const Pi = 3.1415

// Function demonstrating variable usage
func PrintMessage() {
	// Local Variable
	var name string = "Harsh"
	
	// Short Variable Declaration (only inside functions)
	age := 25

	// Multiple Variable Declaration
	var x, y int = 10, 20
	country, city := "India", "Delhi"

	// Default Zero Values
	var emptyString string  // ""
	var emptyInt int        // 0
	var emptyBool bool      // false

	// Print Statements
	fmt.Println("Hello World,", name)
	fmt.Println("Age:", age)
	fmt.Println("Public Variable:", Public)
	fmt.Println("Private Variable:", privateVar)
	fmt.Println("Pi Constant:", Pi)
	fmt.Println("X and Y:", x, y)
	fmt.Println("Country and City:", country, city)
	fmt.Println("Default Values:", emptyString, emptyInt, emptyBool)
}
