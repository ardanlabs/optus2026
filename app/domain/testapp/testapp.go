// Package testapp ....
package testapp

import (
	"context"
	"encoding/json"
	"math/rand/v2"
	"net/http"

	"github.com/ardanlabs/service/app/sdk/errs"
	"github.com/ardanlabs/service/foundation/web"
)

type status struct {
	Status string
}

func (s status) Encode() ([]byte, string, error) {
	data, err := json.Marshal(s)
	return data, "application/json", err
}

func Test(ctx context.Context, w http.ResponseWriter, r *http.Request) web.Encoder {
	if n := rand.IntN(100); n%2 == 0 {
		return errs.Errorf(errs.NotFound, "THIS IS AN ERROR")
	}

	return status{
		Status: "OK",
	}
}
