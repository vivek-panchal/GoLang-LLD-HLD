package main

import "fmt"

func main() {
	// This is the main function where the program starts execution.
	// You can add your code here to implement the desired functionality.

	// Example: Print a message to the console
	fmt.Println("lets start go")

	// ********************************values in go***************************************
	fmt.Println("lets start golang with Vivek") // string
	fmt.Println(42)              // integer
	fmt.Println(3.14)           // float
	fmt.Println(true)          // boolean

	// ********************************variables in go***************************************
	var name string = "Vivek Panchal" // Declare a string variable
	var occupation = "Software Engineer" // Declare a variable with type inference
	var age int = 23            // Declare an integer variable
	var isEmployed = true // Declare a boolean variable
	var experience float64 = 1.5 // Declare a float variable
	var salary = 500000.0 // Declare a variable with type inference
	employer := "Coindcx" // Declare a variable with type inference using short variable declaration { shorthand declaration }

	// Print the values of the variables
	fmt.Println("Name:", name)   // Print the string variable
	fmt.Println("Age:", age)     // Print the integer variable
	fmt.Println("Occupation:", occupation) // Print the variable with type inference
	fmt.Println("Is Employed:", isEmployed) // Print the boolean variable
	fmt.Println("Experience:", experience) // Print the float variable
	fmt.Println("Salary:", salary) // Print the variable with type inference
	fmt.Println("Employer:", employer) // Print the variable with type inference

	// ********************************constants in go***************************************
	// Constants are declared using the `const` keyword and cannot be changed after declaration.
	// They are useful for defining fixed values that should not change during the program execution.

	const pi float64 = 3.14 // Declare a constant for the value of pi
	const gravity = 9.81 // Declare a constant for the value of gravity

	// Declare multiple constants in a single block
	const (
		maxUsers = 100 // Declare a constant for the maximum number of users
		minUsers = 1   // Declare a constant for the minimum number of users
		defaultTimeout = 30 // Declare a constant for the default timeout in seconds
		maxRetries = 5 // Declare a constant for the maximum number of retries
		minRetries = 1 // Declare a constant for the minimum number of retries
		maxConnections = 1000 // Declare a constant for the maximum number of connections
		minConnections = 1 // Declare a constant for the minimum number of connections
		defaultPort = 8080 // Declare a constant for the default port
		defaultHost = "localhost" // Declare a constant for the default host
		defaultProtocol = "http" // Declare a constant for the default protocol
	)
	fmt.Println("Value of Pi:", pi) // Print the value of pi
	fmt.Println("Value of Gravity:", gravity) // Print the value of gravity
	fmt.Println("Max Users:", maxUsers) // Print the maximum number of users
	fmt.Println("Min Users:", minUsers) // Print the minimum number of users
	fmt.Println("Default Timeout:", defaultTimeout) // Print the default timeout
	fmt.Println("Max Retries:", maxRetries) // Print the maximum number of retries
	fmt.Println("Min Retries:", minRetries) // Print the minimum number of retries
	fmt.Println("Max Connections:", maxConnections) // Print the maximum number of connections
	fmt.Println("Min Connections:", minConnections) // Print the minimum number of connections	
	fmt.Println("Default Port:", defaultPort) // Print the default port
	fmt.Println("Default Host:", defaultHost) // Print the default host
	fmt.Println("Default Protocol:", defaultProtocol) // Print the default protocol

	// ********************************for loops in go***************************************
	//only construct in go that can be used to iterate over a range of values or a collection. 
	// The for loop can be used in different forms, including a traditional for loop, a while-like loop, and a range-based loop.
	
	// A while like loop 
	i:= 0 // Initialize a variable i to 0
	for i < 5 { // Loop while i is less than 5
		fmt.Println("Value of i:", i) 
		i++ // Increment i by 1

	}

	//Infinite loop
	/*for {
		fmt.Println("This is an infinite loop") // Print a message in the infinite loop
		// To break out of the infinite loop, you can use a break statement or a condition.
	}*/

	// A traditional for loop
	for j := 0; j < 5; j++ { // Loop from 0
		fmt.Println("Value of j:", j) // Print the value of j
	}

	// A range-based loop
	numbers := []int{1, 2, 3, 4, 5} // Declare a slice of integers
	for index, value := range numbers { // Loop over the slice using range
		fmt.Println("Index:", index, "Value:", value) // Print the index and value
	}

	// ********************************if else in go***************************************
	// The if-else statement is used to execute different blocks of code based on a condition
	if salary > 90000 { // Check if salary is greater than 50000
		fmt.Println("High Salary") // Print a message if the condition is true
	} else if salary > 50000 { // Check if salary is greater than 30000
		fmt.Println("Medium Salary") // Print a message if the condition is true
	} else {
		fmt.Println("Low Salary") // Print a message if the condition is false
	}

	//*********************************switch statement in go***************************************
	switch {
		case salary > 200000: // Check if salary is greater than 200000
			fmt.Println("Very High Salary") // Print a message if the condition is true
		case salary > 90000:
			fmt.Println("High Salary")
		case salary > 50000:
			fmt.Println("Medium Salary")
		default:
			fmt.Println("Low Salary")
	}

	// type switch
	var x interface{} = "Hello" // Declare an interface variable x with a string value
	switch v := x.(type) { // Use a type switch to determine the type of x
	case int:
		fmt.Println("x is an int:", v) // Print if x is an int
	case string:
		fmt.Println("x is a string:", v) // Print if x is a string
	default:
		fmt.Println("x is of unknown type") // Print if x is of an unknown type
	} 

	// ********************************Arrays in go***************************************

	var arr [5]int // Declare an array of integers with a size of 5
	arr[0] = 1 // Assign a value to the first element of the array
	arr[1] = 2 // Assign a value to the second element of the array
	arr[2] = 3 // Assign a value to the third element of the array
	arr[3] = 4 // Assign a value to the fourth element of the array
	arr[4] = 5 // Assign a value to the fifth element of the array
	fmt.Println("Array:", arr) // Print the entire array
	fmt.Println("Length of Array:", len(arr)) // Print the length of the array
	fmt.Println("First Element of Array:", arr[0]) // Print the first element of the array
	fmt.Println("Last Element of Array:", arr[len(arr)-1]) // Print the last element of the array
	fmt.Println("Array Slice:", arr[1:4]) // Print a slice of the array
	fmt.Println("Array Slice with Step:", arr[1:4:5]) // Print a slice of the array with a specified capacity
	fmt.Println("capacity of array is: ", cap(arr)) // Print the capacity of the array

	num := [5]int{1, 2, 3, 4, 5} // Declare and initialize an array of integers
	fmt.Println("Initialized Array:", num) // Print the initialized array
	fmt.Println("Length of Initialized Array:", len(num)) // Print the length of the initialized array

	// 2D Array
	var matrix [3][3]int // Declare a 2D array (matrix) of integers
	matrix[0][0] = 1 // Assign a value to the first element of the first row
	matrix[0][1] = 2 // Assign a value to the second element of the first row
	matrix[0][2] = 3 // Assign a value to the third element of the first row
	matrix[1][0] = 4 // Assign a value to the first element of the second row
	matrix[1][1] = 5 // Assign a value to the second element of the second row
	matrix[1][2] = 6 // Assign a value to the third element of the second row
	matrix[2][0] = 7 // Assign a value to the first element of the
	matrix[2][1] = 8 // Assign a value to the second element of the third row
	matrix[2][2] = 9 // Assign a value to the third element of the third row
	fmt.Println("2D Array (Matrix):", matrix) // Print the entire 2D array

	// *********************************Slices in go***************************************
	// Slices are dynamic arrays that can grow and shrink in size. They are more flexible than arrays and can be used to store collections of data.

	// uninitialized slice
	var uninitializedSlice []int // Declare an uninitialized slice of integers
	fmt.Println("Uninitialized Slice:", uninitializedSlice) // Print the uninitialized slice
	fmt.Println(uninitializedSlice == nil) // Check if the uninitialized slice is nil
	fmt.Println("Length of Uninitialized Slice:", len(uninitializedSlice)) // Print the length of the uninitialized slice

	// initialized slice
	var initializedSlice = make([]int, 5) // Declare and initialize a slice of integers with a length of 5
	initializedSlice[0] = 1 // Assign a value to the first element of the slice
	initializedSlice = append(initializedSlice, 2) // Append a value to the slice
	fmt.Println("Initialized Slice:", initializedSlice) // Print the initialized slice
	fmt.Println("Length of Initialized Slice:", len(initializedSlice)) // Print the length of the initialized slice
	fmt.Println("Capacity of Initialized Slice:", cap(initializedSlice)) // Print the capacity of the initialized slice
	slice := []int{1, 2, 3, 4, 5} // Declare and initialize a slice of integers
	fmt.Println("Slice:", slice) // Print the entire slice
	fmt.Println("Length of Slice:", len(slice)) // Print the length of the slice
	fmt.Println("Capacity of Slice:", cap(slice)) // Print the capacity of the slice

	//2d slice
	twoDSlice := [][]int{ // Declare and initialize a 2D slice of integers
		{1, 2, 3}, // First row
		{4, 5, 6}, // Second row
		{7, 8, 9}, // Third row
	}
	fmt.Println("2D Slice:", twoDSlice) // Print the entire 2D slice
	fmt.Println("Length of 2D Slice:", len(twoDSlice)) // Print the length
	fmt.Println("Length of First Row of 2D Slice:", len(twoDSlice[0])) // Print the length of the first row of the 2D slice
	fmt.Println("Capacity of 2D Slice:", cap(twoDSlice)) // Print the capacity of the 2D slice
	fmt.Println("Capacity of First Row of 2D Slice:", cap(twoDSlice[0])) // Print the capacity of the first row of the 2D slice

	// ********************************Maps in go***************************************
	// Maps are unordered collections of key-value pairs. They are similar to dictionaries in other programming languages and can be used to store data in a way that allows for fast lookups.

	// uninitialized map
	var uninitializedMap map[string]int // Declare an uninitialized map with string keys and integer values
    // setting elements
	uninitializedMap = make(map[string]int) // Initialize the map using make
	uninitializedMap["apple"] = 1 // Set a key-value pair in the map
	uninitializedMap["banana"] = 2 // Set another key-value pair in the map
	uninitializedMap["orange"] = 3 // Set another key-value pair in the map

	// Print the map and check if it is nil
	fmt.Println("Uninitialized Map:", uninitializedMap) // Print the uninitialized map
	fmt.Println(uninitializedMap == nil) // Check if the uninitialized map is nil
	fmt.Println("Length of Uninitialized Map:", len(uninitializedMap)) // Print the length of the uninitialized map
	fmt.Println("Uninitialized Map:", uninitializedMap) 
	fmt.Println("Value for key 'apple':", uninitializedMap["apple"]) // Print the value for the key "apple"
	fmt.Println("Value for key 'banana':", uninitializedMap["banana"]) // Print the
	// value for the key "banana"
	fmt.Println("Value for key 'orange':", uninitializedMap["orange"]) // Print the value for the key "orange"

	// delete a key-value pair from the map
	delete(uninitializedMap, "banana") // Delete the key-value pair with the key "banana" from the map
	fmt.Println("Map after deleting 'banana':", uninitializedMap) // Print the map after deleting "banana"

	// clear the map using clear function
	clear(uninitializedMap) // Clear the map using the clear function
	fmt.Println("Map after clearing:", uninitializedMap) // Print the map after clearing it

	//**********************************Range in go***************************************
	// The range keyword is used to iterate over elements in a collection, such as arrays, slices, maps, and strings.
	// It can be used in for loops to access both the index and value of each element
	var numbersMap = map[string]int{"one": 1, "two": 2, "three": 3} // Declare and initialize a map with string keys and integer values
	fmt.Println("Numbers Map:", numbersMap) // Print the entire map
	for key, value := range numbersMap {
		fmt.Println("Key:", key, ", Value:", value)
	}

}