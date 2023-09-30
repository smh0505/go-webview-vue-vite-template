# Go Webview Template with Vite and Vue

## From Scratch

1. Create your `backend` folder.

```cmd
mkdir backend
cd backend
go mod init your-project-name
go get github.com/webview/webview_go
```

2. Create your `frontend` folder using vite.

```cmd
yarn create vite frontend

:: or

npm create vite@latest frontend
```

3. Write `main.go`, `dev.go`, and `release.go`.

```go
// main.go
package main

import webview "github.com/webview/webview_go"

var w webview.WebView

func main() {
    // Create window
    debug := true
    w = webview.New(debug)
    defer w.Destroy()

    // Set default size and title of the window
    w.SetSize(800, 600, webview.HintNone)
    w.SetTitle("your-project-name")

    // Connect to the server
    connect()

    // Bind functions
    w.Bind("js_function", go_function)
    ...

    // Begin process
    w.Run()
}
```

```go
// dev.go
// +build !release <= NECESSARY

package main

func connect() {
    // Follow the port number of vite dev server
    w.Navigate("http://localhost:5173")
}
```

```go
// release.go
// +build release <= NECESSARY

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
```

4. Start the dev server and compile the program.

```cmd
:: dev.bat
start cmd /c go run ./backend
yarn --cmd ./frontend dev
```

5. Build both frontend and backend into `build` folder

```js
// vite.config.js
export default defineConfig({
    ...,
    build: {
        outDir: "../build/dist",
        emptyOutDir: true
    }
})
```

```cmd
:: build.bat
rmdir /s /q "./build"
mkdir build
go build -C ./backend -o ../build -ldflags -H=windowsgui -tags release
yarn --cmd ./frontend build
```
