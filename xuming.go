package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

func main() {
	prepareQuotes()
	http.HandleFunc("/", redirectHandler)
	http.HandleFunc("/+1s", plus1sHandler)
	http.HandleFunc("/+1s/beat", plus1sHeartbeatHandler)
	http.HandleFunc("/+1s/quote", quoteHandler)
	log.Fatal(http.ListenAndServe("localhost:8080",  nil))
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "+1s", 302)
}

var life uint64
var online int64
var tmpl, _ = template.ParseFiles("template/index.html")

func plus1sHandler(w http.ResponseWriter, r *http.Request) {
	current := atomic.AddUint64(&life, 1)
	tmpl.Execute(w, struct {
		Life uint64
		Online int64
		Quote string
	} {
		Life: current,
		Online: online,
		Quote: getQuote(),
	})
}

func plus1sHeartbeatHandler(w http.ResponseWriter, r *http.Request) {
	current := atomic.AddUint64(&life, 1)
	w.Header().Set("Content-Type", "application/json; encoding=UTF-8")
	fmt.Fprintf(w, "{\"life\":%d,\"online\":%d}", current, online)
}

func quoteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "test/plain; encoding=UTF-8")
	w.WriteHeader(200)
	atomic.AddInt64(&online, 1)
	time.Sleep(9700 * time.Millisecond)
	atomic.AddInt64(&online, -1)
	w.Write([]byte(getQuote()))
}
