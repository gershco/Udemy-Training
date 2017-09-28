package main
import ("fmt"
"bufio"
"os")


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

// Creates a function prod which multiplies two integers
func prod(e int, f int) int {
	return e * f
}


func main() {

for {

  fmt.Print(`Which calculation do you want to run?

  1) Addition
  2) Subtraction
  3) Multiplication
  4) None; I've had enough!

Please enter 1..4 and press ENTER: `)

reader := bufio.NewReader(os.Stdin)
result, _, err := reader.ReadRune()
if err != nil {
	fmt.Println(err)
	return
}

switch result {

case '1':

  // Enter first number and set to x
      fmt.Println("Enter your first number: ")
      var x int
      fmt.Scanln(&x)

  // Enter second number and set to y
      fmt.Println("Enter your second number: ")
      var y int
      fmt.Scanln(&y)

  //Adds the two numbers using the add function below
      fmt.Println("The sum of your numbers is",add(x,y),".")
 
case '2':

  // Enter first number and set to x
      fmt.Println("Enter your first number: ")
      var x int
      fmt.Scanln(&x)

  // Enter second number and set to y
      fmt.Println("Enter your second number: ")
      var y int
      fmt.Scanln(&y)

  //Subtracts the two numbers using the minus function below
      fmt.Println("The difference between your numbers is",minus(x,y),".")

case '3':

  // Enter first number and set to x
      fmt.Println("Enter your first number: ")
      var x int
      fmt.Scanln(&x)

  // Enter second number and set to y
      fmt.Println("Enter your second number: ")
      var y int
      fmt.Scanln(&y)

  //Multiplies the two numbers using the prod function below
      fmt.Println("The difference between your numbers is",prod(x,y),".")

case '4':
  fmt.Println("Oh, that's a shame!")
  os.Exit(0)
 
default:
}}}