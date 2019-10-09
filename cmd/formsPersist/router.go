package main

import (
	"net/http"

	"github.com/MrMikeandike/FormsPersistGo/pkg/webutils"
)

// App is main router for the web app. it is the entrypoint
type App struct {
	InputHandler *InputHandler
}

func (h *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = webutils.ShiftPath(r.URL.Path)
	switch head {
	case "input":
		h.InputHandler.ServeHTTP(w, r)
	default:
		NotFound(w)
		return
	}
	return
}

// InputHandler handles both the gets for the form and the posts
// path: begins with /input
type InputHandler struct {
}

func (*InputHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = webutils.ShiftPath(r.URL.Path)
	switch head {
	case "":
		switch r.Method {
		case http.MethodGet:

		case http.MethodPost:
		default:
			MethodNotAllowed(w)
			return

		}
	default:
		NotFound(w)
		return
	}
	return
}

// Get method of inputhandler handles get method requests to /input
func (*InputHandler) Get() http.Handler {
	// code if you want
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		return
	})
}

// NotFound sends not found error to w
func NotFound(w http.ResponseWriter) {
	http.Error(w, "404 not found!!!", http.StatusNotFound)
}

// MethodNotAllowed sends MethodNotAllowed error to w
func MethodNotAllowed(w http.ResponseWriter) {
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
