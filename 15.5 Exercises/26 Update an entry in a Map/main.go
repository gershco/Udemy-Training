package main

import "fmt"

func main() {

	dictionary := map[string]string{

		"apple": " A rounded fruit with a red, yellow, or green skin, firm white flesh and a seedy core.",
		"boat":  " A small vessel for travel on water.",
		"cargo": " The goods transported in a ship, airplane, or vehicle.",
	}
	fmt.Println()
	fmt.Println(dictionary)
	fmt.Println("The map has a length of", len(dictionary))

	dictionary["dart"] = " A small pointed object that is meant to be thrown."

	fmt.Println()
	fmt.Println(dictionary)

	fmt.Println()
	fmt.Println("After adding another entry to the map, the length is now ", len(dictionary))

	fmt.Println()
	fmt.Println("The definition of boat is:", dictionary["boat"])
}
