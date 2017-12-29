package main

import (
	"io"
	"net/http"
)

//Entry is used by different functions.
//var Entry string

func main() {
	http.HandleFunc("/grab", grab)
	http.HandleFunc("/", index)

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8081", nil)

}

func grab(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		http.ServeFile(w, r, "grab.html")

	} else {
	}
}

/*

func main() {
	// Load a CSV file.
	f, _ := os.Open("crime.csv")

	// Create a new reader.
	r := csv.NewReader(bufio.NewReader(f))
	for {
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}
		// Display record.
		// ... Display record length.
		// ... Display all individual elements of the slice.
		fmt.Println(record)
		fmt.Println(len(record))
		for value := range record {
			fmt.Printf("  %v\n", record[value])
		}
	}
}

*/

func index(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, "<b>You need to browse to localhost:8081/grab.<b>")
}
