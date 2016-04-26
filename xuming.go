package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

var life uint64
var online uint64

func main() {
	prepareQuotes()
	http.HandleFunc("/", redirectHandler)
	http.HandleFunc("/+1s", plus1sHandler)
	http.HandleFunc("/+1s/beat", plus1sHeartbeatHandler)
	http.HandleFunc("/+1s/quote", quoteHandler)
	log.Fatal(http.ListenAndServe(":8080",  nil))
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "+1s", 302)
}

func plus1sHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/index.html")
	t.Execute(w, struct {
		Life uint64
		Online uint64
		Quote string
	} {
		Life: current,
		Online: lastonline+1,
		Quote: getQuote(),
	})
}

func plus1sHeartbeatHandler(w http.ResponseWriter, r *http.Request) {
	current := atomic.AddUint64(&life, 1)
	w.Header().Set("Content-Type", "application/json; encoding=UTF-8")
	fmt.Fprintf(w, "{\"life\":%d,\"online\":%d}", current, lastonline)
}

func quoteHandler(w http.ResponseWriter, r *http.Request) {
	atomic.AddUint64(&online, 1)
	time.Sleep(9700 * time.Milisecond)
	atomic.AddUint64(&online, -1)
	w.Header().Set("Content-Type", "test/plain; encoding=UTF-8")
	w.Write([]byte(getQuote()))
}
