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
var lastonline uint64

func main() {
	go updateOnline()
	http.HandleFunc("/", redirectHandler)
	http.HandleFunc("/+1s", plus1sHandler)
	http.HandleFunc("/+1s/beat", plus1sHeartbeatHandler)
	log.Fatal(http.ListenAndServe(":8080",  nil))
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "+1s", 302)
}

func plus1sHandler(w http.ResponseWriter, r *http.Request) {
	current := atomic.AddUint64(&life, 1)
	atomic.AddUint64(&online, 1)
	t, _ := template.ParseFiles("template/index.html")
	t.Execute(w, struct {
		Life uint64
		Online uint64
	} {
		Life: current,
		Online: lastonline+1,
	})
}

func plus1sHeartbeatHandler(w http.ResponseWriter, r *http.Request) {
	current := atomic.AddUint64(&life, 1)
	atomic.AddUint64(&online, 1)
	fmt.Fprintf(w, "{\"life\":%d,\"online\":%d}", current, lastonline)
}

func updateOnline() {
	for {
		time.Sleep(time.Second)
		lastonline, online = online, 0
	}
}
