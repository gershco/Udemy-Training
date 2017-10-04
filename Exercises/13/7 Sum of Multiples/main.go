package main

import "fmt"

func main () {

	var s int

	for i := 1 ; i < 1000 ; i++ {

		if i%3 == 0 || i%5 == 0 {

			s+=i			
			
		} 		
		
	}

	fmt.Println("The sum of all the multiples of 3 or 5 below 1000 is" , s)
}