package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil{
		fmt.Fprintf(w, "ParseForm() err %v", err)
	}
	fmt.Fprintf(w, "POST request successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "NAME = %s\n", name)
	fmt.Fprintf(w, "ADDRESS= %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET"{
		http.Error(w, "method is not supported", http.StatusNotFound)

	}

	fmt.Fprintf(w, "hello!")
}

func main(){
	// to check out the static directory 
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	//
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello",helloHandler)

	fmt.Printf("stating server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil{
		log.Fatal(err)
	}
}