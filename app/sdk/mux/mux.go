// Package mux ...
package mux

import (
	"github.com/ardanlabs/service/app/domain/testapp"
	"github.com/ardanlabs/service/foundation/web"
)

func WebAPI() *web.App {
	app := web.NewApp()

	app.HandleFunc("/test", testapp.Test)

	return app
}
