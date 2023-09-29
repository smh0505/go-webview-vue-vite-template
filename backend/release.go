// +build release

package main

import "net/http"

func connect() {
	go start()
	w.Navigate("http://localhost:3000")
}

func start() {
	fs := http.FileServer(http.Dir("./dist"))
	http.Handle("/", fs)
	http.ListenAndServe(":3000", nil)
}
