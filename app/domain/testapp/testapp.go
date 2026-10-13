// Package testapp ....
package testapp

import (
	"context"
	"encoding/json"
	"net/http"
)

func Test(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	// DECODE INPUT (JSON)
	// CALL BUSINESS LAYER
	// CREATE RESPONSE DATA

	status := struct {
		Status string
	}{
		Status: "OK",
	}

	return json.NewEncoder(w).Encode(status)
}
