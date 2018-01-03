package main

import (
	"encoding/json"
	"fmt"

	_ "github.com/lib/pq"
)

type XMan struct {
	Data []Man `json:"data"`
}

type Man struct {
	ID        int
	Age       int
	FirstName string
	LastName  string
	Email     string
	Balance   float64
}

var men XMan

func main() {

	fmt.Printf("men is of type %T\n\n", men)

	m1 := Man{1, 25, "John", "Minor", "john@minor.gov", 125.38}
	m2 := Man{2, 34, "James", "Riley", "james@riley.com", 78.56}
	m3 := Man{3, 46, "Nigel", "Kennedy", "nigel@kennedy.mus", 21.59}

	men := []XMan{m1, m2, m3}

	fmt.Printf("\nmen is of type %T\n\n", men)

	list, _ := json.Marshal(men)

	fmt.Println(string(list))
}
