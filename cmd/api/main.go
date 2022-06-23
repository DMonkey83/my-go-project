package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	serveMux := http.NewServeMux()

	serveMux.HandleFunc("/", testHandler)

	const addr = "localhost:8080"
	srv := http.Server{
		Handler:      serveMux,
		Addr:         addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	fmt.Println("server started on ", addr)
	fmt.Println("server started on ", addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}

func testHandler(w http.ResponsEWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte("{}"))
}
