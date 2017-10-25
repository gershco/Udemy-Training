package main

import "fmt"

func main() {

	dictionary := make(map[string]string)

	dictionary["apple"] = "a rounded fruit with a red, yellow, or green skin, firm white flesh and a seedy core"
	dictionary["boat"] = "a small vessel for travel on water"
	dictionary["cargo"] = "the goods transported in a ship, airplane, or vehicle"

	fmt.Println(len(dictionary))
}
