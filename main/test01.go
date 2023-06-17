package main

import (
	"fmt"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

//func main() {
//	http.HandleFunc("/", handler)
//	http.HandleFunc("/count", counter)
//	imgHandle := func(w http.ResponseWriter, r *http.Request) {
//		lissajous(w)
//	}
//	http.HandleFunc("/img", imgHandle)
//	log.Fatal(http.ListenAndServe(":8081", nil))
//}

func printCurrentUrl(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock() //加锁
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
