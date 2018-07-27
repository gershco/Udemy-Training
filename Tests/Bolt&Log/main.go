package main

import (
	"fmt"
	"os"
	"path/filepath"
)

//CheckOrCreateFolder will check if folderName exists in base folder otherwise create new
func CheckOrCreateFolder(folderName string) {
	//fileabs
	dir, errpath := filepath.Abs(filepath.Dir(os.Args[0]))
	if errpath != nil {
		fmt.Println("Directory not retrieved")
	}
	fmt.Println(dir)

	if _, err := os.Stat(dir + "/" + folderName); err == nil {
		// /myfolder exists
		fmt.Println(dir + "/" + folderName + " exists")
	} else {
		// the path does not exist or some error occurred.
		fmt.Println(dir + "/" + folderName + " does not exist")
		err := os.Mkdir(dir+"/"+folderName, 0777)
		if err != nil {
			fmt.Println("could not create" + dir + "/" + folderName)
		} else {
			fmt.Println(dir + "/" + folderName + " created")
		}

	}
}
func main() {

	CheckOrCreateFolder("myfolder")
}
