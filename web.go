package main

import (
	"net/http"
	"os/exec"
	"runtime"
)

// WebServer represents a simple web server
type WebServer struct {
	static http.Handler
}

// NewWebServer creates a new WebServer instance
func NewWebServer(dir string) (*WebServer, error) {
	ws := &WebServer{}

	s := http.FileServer(http.Dir(dir))
	ws.static = s
	return ws, nil
}

// Bind starts the web server listening on the given address / port
func (ws *WebServer) Bind(bind string) {
	http.ListenAndServe(bind, ws.static)
}

// Run starts th web server listening on port 8080
func (ws *WebServer) Run() {
	ws.Bind(":8080")
}

func (ws *WebServer) OpenBrowser() {
	var command string
	var args []string

	switch runtime.GOOS {
	case "windows":
		command = "cmd"
		args = []string{"/c", "start"}
	case "osx":
		command = "open"
	default:
		command = "xdg-open"
	}

	args = append(args, "http://127.0.0.1:8080")

	err := exec.Command(command, args...).Start()
	if err != nil {
		panic(err)
	}
}
