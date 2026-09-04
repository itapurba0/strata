package main

import (
	"fmt"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w,"STARTA API is Live and Healthy...")
}

func main(){
	http.HandleFunc("/health",healthHandler)
	fmt.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}