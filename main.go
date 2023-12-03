package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	log.Println("starting app")
	http.HandleFunc("/", greet)
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal("failed to launch app", err)
	}
	log.Println("terminationg app")
}
