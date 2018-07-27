package main

import (
	"fmt"
	"os"
	"path"
	"runtime"
)

func main() {

	_, fullPath, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("Directory not retrieved")
	}
	fmt.Println(fullPath)
	trimmedPath := path.Dir(fullPath)
	fmt.Println(trimmedPath)

	if _, err := os.Stat(trimmedPath + "/myfolder"); err == nil {

		// /myfolder exists
		fmt.Println(trimmedPath + "/myfolder exists")
	} else {
		// the path does not exist or some error occurred.
		fmt.Println(trimmedPath + "/myfolder does not exist")
		err := os.Mkdir(trimmedPath+"/myfolder", 0777)
		if err != nil {
			fmt.Println("could not create" + trimmedPath + "/myfolder")
		} else {
			fmt.Println(trimmedPath + "/myfolder created")
		}

	}
}
