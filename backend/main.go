package main

import (
	"github.com/pkg/browser"
	webview "github.com/webview/webview_go"
)

var w webview.WebView

func main() {
	// Create window
	debug := true
	w = webview.New(debug)
	defer w.Destroy()

	// Set default size of the window
	w.SetSize(800, 600, webview.HintNone)
	w.SetTitle("Go Webview Template with Vite and Vue")

	// Connect to the server
	connect()

	// Bind functions
	w.Bind("openAnotherBrowser", openLink)

	// Begin process
	w.Run()
}

func openLink(link string) error {
	return browser.OpenURL(link)
}
