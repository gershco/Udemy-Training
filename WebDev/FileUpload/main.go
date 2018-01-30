package main

import (
	"fmt"
	"html/template"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/nfnt/resize"
)

// compile all templates and cache them
//var templates = template.Must(template.ParseGlob("templates/*"))

func upload(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("index.html")

	t.Execute(w, nil)

}

func ajaxReturn(w http.ResponseWriter, r *http.Request) {

	// Post
	fmt.Println("Starting to process file")
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	f, err := os.OpenFile("./uploadedFiles/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	io.Copy(f, file)

	// open "test.jpg"
	uploadedFile, err := os.Open("./uploadedFiles/" + handler.Filename)
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(uploadedFile)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// resize to thumbnail using Lanczos resampling
	// and preserve aspect ratio
	resizedImage := resize.Thumbnail(75, 75, img, resize.Lanczos3)

	out, err := os.Create("./resizedImages/resized" + handler.Filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, resizedImage, nil)

	fmt.Fprintln(w, "Your file has been uploaded and resized.")

}

func main() {

	fmt.Println("Serving at localhost:8080/upload")
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/return", ajaxReturn)
	http.ListenAndServe(":8080", nil) // setting listening port
}
