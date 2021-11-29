package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Go HTTP")
	})

	http.HandleFunc("/time", func(rw http.ResponseWriter, r *http.Request) {
		t := time.Now()
		timeString := fmt.Sprintf("{\"time\": \"%s\"}", t)
		rw.Write([]byte(timeString))
	})
	http.ListenAndServe(":8081", nil)
}
