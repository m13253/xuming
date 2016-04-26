package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sync/atomic"
)

var life uint64

func main() {
	http.HandleFunc("/", redirectHandler)
	http.HandleFunc("/+1s", plus1sHandler)
	http.HandleFunc("/+1s/beat", plus1sHartbeatHandler)
	log.Fatal(http.ListenAndServe(":8080",  nil))
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "+1s", 302)
}

func plus1sHandler(w http.ResponseWriter, r *http.Request) {
	current := atomic.AddUint64(&life, 1)
	t, _ := template.ParseFiles("template/index.html")
	t.Execute(w, struct {
		Life uint64
	} {
		Life: current,
	})
}

func plus1sHartbeatHandler(w http.ResponseWriter, r *http.Request) {
	current := atomic.AddUint64(&life, 1)
	fmt.Fprintf(w, "%d", current)
}
