package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// instead of using this function which would be hard to test because it writes strictly to Stdout,
// we use the below one instead, to increase the testability of the function
// func Greet(writer *bytes.Buffer, name string) {
// 	fmt.Printf("Hello, %s", name)
// }

// generalised function that writes hello,<INPUT> into a writer
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// Greet can be used to take in a ResponseWriter that can write into a HTTP response as a handler 
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	// Greet can be used to take in a os.Stout and write into the terminal 
	Greet(os.Stdout, "Elodie")
	
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}

