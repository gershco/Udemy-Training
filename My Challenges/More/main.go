package main

import ("fmt"
    "math")

// Enter first number and set to x
func main() {
    fmt.Println("Enter your first number: ")
    var x int
    fmt.Scanln(&x)

// Enter second number and set to y
    fmt.Println("Enter your second number: ")
    var y int
    fmt.Scanln(&y)

//Adds the two numbers using the add function below
    fmt.Println("The sum of your numbers is",add(x,y),".")

//Subtracts the two numbers using the minus function below
    fmt.Println("The difference between your numbers is",minus(x,y),".")

    //Squares the numbers

        fmt.Println("The squares of your numbers are", sq(x), "and ", sq(y))
    }


// Creates a function add which sums two integers
func add(a int, b int) int {
	return a + b
}
// Creates a function minus which calculates the positive difference between two integers
func minus(c int, d int) int {
    if c>d {
    return c-d
    } else {
    return d-c
}
}


// Creates a function to square a given integer
func sq (e int) float64 {
g := float64(e)
    return math.Pow(g,2)
}
