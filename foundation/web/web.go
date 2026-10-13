// Package web ...
package web

import (
	"context"
	"fmt"
	"net/http"
)

type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

type App struct {
	*http.ServeMux
}

func NewApp() *App {
	return &App{
		ServeMux: http.NewServeMux(),
	}
}

// HandleFunc is my new handler func.
func (app *App) HandleFunc(pattern string, handler Handler) {

	h := func(w http.ResponseWriter, r *http.Request) {

		// I CAN PUT ANY CODE HERE

		if err := handler(r.Context(), w, r); err != nil {
			fmt.Println(err)
			return
		}

		// I CAN PUT ANY CODE HERE
	}

	app.ServeMux.HandleFunc(pattern, h)
}
