package main

import (
	"fmt"
	"log"
	"net/http"
)

// ab testing
// ab -c 100 -n 5000 http://127.0.0.1:8080/

func main() {
	// Limit to 1 CPU
	//runtime.GOMAXPROCS(1)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, world!")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
