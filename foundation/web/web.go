// Package web ...
package web

import (
	"context"
	"fmt"
	"net/http"
)

type HandlerFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

type App struct {
	*http.ServeMux
	mw []MidFunc
}

func NewApp(mw ...MidFunc) *App {
	return &App{
		ServeMux: http.NewServeMux(),
		mw:       mw,
	}
}

// HandleFunc is my new handler func.
func (a *App) HandleFunc(pattern string, handlerFunc HandlerFunc, mw ...MidFunc) {
	handlerFunc = wrapMiddleware(mw, handlerFunc)
	handlerFunc = wrapMiddleware(a.mw, handlerFunc)

	h := func(w http.ResponseWriter, r *http.Request) {

		// I CAN PUT ANY CODE HERE

		if err := handlerFunc(r.Context(), w, r); err != nil {
			fmt.Println(err)
			return
		}

		// I CAN PUT ANY CODE HERE
	}

	a.ServeMux.HandleFunc(pattern, h)
}
