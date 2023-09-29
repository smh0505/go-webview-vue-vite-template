package main

import webview "github.com/webview/webview_go"

var w webview.WebView

func main() {
	debug := true
	w = webview.New(debug)
	defer w.Destroy()

	w.SetSize(800, 600, webview.HintNone)
	w.SetTitle("Go Webview Template with Vite and Vue")

	connect()
	w.Run()
}
