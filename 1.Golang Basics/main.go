package main

import (
	"fmt"
	"time"
	"sync"
)

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

	//***********************************Functions in go***************************************
	// Functions are reusable blocks of code that can be called with specific parameters to perform a task
	// Functions can return values and can be defined with parameters to accept input values.
	// Call the function to print "Hello, World!"
	printHello()

	// Call the function to add two numbers and print the result
	result := add(5, 10) // Call the add function with parameters 5 and 10
	fmt.Println("Sum of 5 and 10 is:", result) // Print the result of the addition

	// calling function with multiple return values
	sum, product := addMultiple(5, 10) // Call the addMultiple function with parameters 5 and 10
	fmt.Println("Sum of 5 and 10 is:", sum) // Print the sum returned
	fmt.Println("Product of 5 and 10 is:", product) // Print the product returned

	// Call the applyFunction function with the add function and parameters 5 and 10
	applyResult := applyFunction(add, 5, 10) // Call the applyFunction function with the add function and parameters 5 and 10
	fmt.Println("Result of applying add function:", applyResult) // Print the result of applying the

	// Variadic functions :- We can pass a variable number of parameters to a function using the variadic syntax.
	// The variadic function can accept zero or more parameters of a specified type.
	sumVariadic := sumVariadic(1, 2, 3, 4, 5, 6, 7, 8) // Call the sumVariadic function with multiple parameters
	fmt.Println("Sum of variadic parameters:", sumVariadic) // Print the result of the sumVariadic function

	//**********************************Closures in go***************************************
	// Closures are functions that capture variables from their surrounding context. They can be used to create functions with specific behavior based on the captured variables.
	// A closure is a function that captures variables from its surrounding context.
	incrementer := makeIncrementer() // Create a closure using the makeIncrementer function
	fmt.Println("Incrementer with 5:", incrementer(5)) // Call the closure with the parameter 5
	fmt.Println("Incrementer with 10:", incrementer(10)) // Call the closure with the parameter 10


	//**********************************Pointers in go***************************************
	// Pointers are variables that store the memory address of another variable. They allow you to manipulate the value of a variable directly.
	var temp int = 10 // Declare an integer variable temp and assign it a value of 10
	var ptr *int = &temp // Declare a pointer variable ptr and assign it the address of temp using the address-of operator (&)
	fmt.Println("Value of temp:", temp) // Print the value of temp
	fmt.Println("Address of temp:", &temp) // Print the address of temp
	fmt.Println("Value of ptr:", ptr) // Print the value of ptr (address of temp)	

	//***********************************Structs in go***************************************
	// Structs are composite data types that allow you to group related data together. They can be used to define custom data structures with fields of different types.
	

	// Create an instance of the Order struct using the newOrder function
	order1 := newOrder(1, "Laptop", 2, 150000.00) // Call the newOrder function to create a new order
	fmt.Println("Order 1:", order1) // Print the details of the first order

	// Create an instance of the Order struct and initialize its fields
	order := Order{
		ID:       1, // Set the order ID
		Product:  "Laptop", // Set the product name	
		Quantity: 2, // Set the quantity of the product
		Price:    150000.00, // Set the price of the product
		//Total:    300000.00, // Set the total price (Quantity * Price)
		createdAt: time.Now(), // Set the creation time of the order to the current time
	}

	// Calculate the total price of the order
	order.Total = float64(order.Quantity) * order.Price // Calculate the total price (Quantity * Price)

	// Print the details of the order
	fmt.Println("Order ID:", order.ID) // Print the order ID
	fmt.Println("Product:", order.Product) // Print the product name
	fmt.Println("Quantity:", order.Quantity) // Print the quantity of the product	
	fmt.Println("Price:", order.Price) // Print the price of the product
	fmt.Println("Total:", order.Total) // Print the total price (Quantity * Price)
	fmt.Println("Order Created At:", order.createdAt) // Print the creation time of the order

	myOrder := Order{ // Create another instance of the Order struct
		ID:       2, // Set the order ID
		Product:  "Smartphone", // Set the product name
		Quantity: 4, // Set the quantity of the product
		Price:    50000.00, // Set the price of the product
		Total:    200000.00, // Set the total price (Quantity * Price)
		createdAt: time.Now(), // Set the creation time of the order to the current time
	}

	// Print the details of the second order
	fmt.Println("myOrder:", myOrder) // Print the order
	fmt.Println("myOrder ID:", myOrder.ID) // Print the order ID
	fmt.Println("myOrder Product:", myOrder.Product) // Print the product name
	fmt.Println("myOrder Quantity:", myOrder.Quantity) // Print the quantity of the product
	fmt.Println("myOrder Price:", myOrder.Price) // Print the price of the product
	fmt.Println("myOrder Total:", myOrder.Total) // Print the total price (Quantity * Price)
	fmt.Println("myOrder Created At:", myOrder.createdAt.Format(time.RFC3339)) // Print the creation time of the order in RFC3339 format



	// *********************************Interfaces in go***************************************
	// Interfaces are abstract types that define a set of methods that a type must implement. They allow you to define behavior without specifying the concrete implementation.
	// Interfaces are used to define a contract that types must adhere to. They allow for polymorphism and can be used to create flexible and reusable code.
	// An interface is a type that defines a set of methods that a type must implement.
	type Shape interface { // Define an interface named Shape
		Area() float64 // Method to calculate the area of the shape
		Perimeter() float64 // Method to calculate the perimeter of the shape
	}

	// Create an instance of the Rectangle type
	rect := Rectangle{ // Create a new instance of the Rectangle type
		Width:  10,
		Height: 5,
	}
	// Print the area and perimeter of the rectangle
	fmt.Println("Rectangle Area:", rect.Area()) // Call the Area method of the Rectangle type
	fmt.Println("Rectangle Perimeter:", rect.Perimeter()) // Call the Perimeter method of the Rectangle type

	//***********************************Enums in go***************************************

	// Enums are a way to define a set of named constants. They can be used to represent a fixed set of values, such as days of the week, months of the year, etc.

	// ***********************************generics in go***************************************
	// Generics allow you to write functions and types that can work with any data type. They provide a way to create reusable code that can operate on different types without sacrificing type safety.

	// Comparable types can be used with generics to create functions that can work with any type that supports comparison operations.


	// Call the generic function with different types
	Print("Hello, Generics!") // Call the Print function with a string value	
	Print(42) // Call the Print function with an integer value
	Print(3.14) // Call the Print function with a float value
	Print(true) // Call the Print function with a boolean value

	// **********************************goroutines in go***************************************
	// Goroutines are lightweight threads of execution that allow you to run concurrent tasks in Go.
	// They are used to perform tasks concurrently, allowing for efficient use of system resources and improved performance.
	fmt.Println("Starting Goroutine") // Print a message before starting the goroutine
	goroutineExample() // Call the goroutineExample function to start a goroutine

	//**********************************Wait Groups in go***************************************
	// Wait Groups are used to wait for a collection of goroutines to finish executing.
	// They provide a way to synchronize the completion of multiple goroutines.
	
	var wg sync.WaitGroup // Declare a WaitGroup variable
	wg.Add(2) // Add two goroutines to the WaitGroup

	go func() {
		defer wg.Done() // Decrement the counter when the goroutine completes
		printHello()
	}()

	go func() {
		defer wg.Done() // Decrement the counter when the goroutine completes
		printHello()
	}()

	wg.Wait() // Wait for all goroutines to complete
}

// Function to print a message
func printHello() {
	fmt.Println("Hello, World!") // Print a message to the console
}

// Function with parameters and return value
func add(a int, b int) int {
	return a + b // Return the sum of a and b
}

// Function with multiple return values
func addMultiple(a int, b int) (int, int) {
	return a + b, a * b // Return the sum and product of a and b
}

// Funtion takes a function as an argument and returns a function
func applyFunction(fn func(int, int) int, a int, b int) int {
	return fn(a, b) // Call the function fn with parameters a and b and return the result
}

// Variadic function that takes a variable number of parameters
func sumVariadic(nums ...int) int {		
	sum := 0 // Initialize a variable to store the sum
	for _, num := range nums { // Iterate over the variadic parameters
		sum += num // Add each number to the sum
	}
	return sum // Return the total sum
}

// Closure example
func makeIncrementer() func(int) int {
	count := 2
	return func(x int) int {
		count += x
		return count
	}
}

// Structs in Go
type Order struct { // Define a struct named Order
	ID       int    // Field for the order ID
	Product  string // Field for the product name
	Quantity int    // Field for the quantity of the product
	Price    float64 // Field for the price of the product
	Total    float64 // Field for the total price (Quantity * Price)
	createdAt time.Time // Field for the creation time of the order
}

// Function to create a new order
// This function takes parameters to create a new order and returns an instance of the Order struct.
func newOrder(id int, product string, quantity int, price float64) Order {
		
	return Order{
		ID:       id,
		Product:  product,
		Quantity: quantity,
		Price:    price,
		Total:    float64(quantity) * price,
		createdAt: time.Now(),
	}
}

// A type can implement an interface by defining the methods specified in the interface.
type Rectangle struct {
	Width  float64
	Height float64
}
// Implement the Area method for the Rectangle type
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Implement the Perimeter method for the Rectangle type
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}


// Example of a generic function
func Print[T any](value T) {
	fmt.Println(value)
}

// Goroutine example
func goroutineExample() {
	go func() { // Start a new goroutine
		fmt.Println("This is a goroutine") // Print a message in the goroutine
	}()

	time.Sleep(1 * time.Second) // Sleep for 1 second to allow the goroutine to complete
	fmt.Println("Goroutine finished") // Print a message after the goroutine has finished
}