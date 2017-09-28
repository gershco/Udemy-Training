package main
import ("fmt"
"bufio"
"os")




func main() {


fmt.Println("Please enter a vowel.")

    reader := bufio.NewReader(os.Stdin)
    result, _, err := reader.ReadRune()
    
    if err != nil {
      fmt.Println(err)
    }
    
    
    switch result {

    case 'a':
      fmt.Println("A is for apple.")
      break

    case 'e':
      fmt.Println("E is for elephant")
      break

      default :
      fmt.Println("You didn't enter a vowel.")


    }

}