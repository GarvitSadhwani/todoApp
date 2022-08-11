package main

import (
	"fmt"
	"log"
	"net/http"
)

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("task saved successfully\n")
	name := r.FormValue("name")
	details := r.FormValue("details")
	fmt.Printf("Submitted details: \nTask name: %s \nTask details: %s\n", name, details)

}

func main() {
	fileServerIndex := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServerIndex)
	http.HandleFunc("/addTask", addTaskHandler)
	fmt.Println("Server started at 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
