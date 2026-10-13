package mid

import (
	"context"
	"net/http"

	"github.com/ardanlabs/kronk/cmd/server/foundation/logger"
	"github.com/ardanlabs/service/foundation/web"
)

func Errors(log *logger.Logger) web.MidFunc {
	m := func(handler web.HandlerFunc) web.HandlerFunc {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) web.Encoder {
			resp := handler(ctx, w, r)
			// if err != nil {
			// 	log.Info(ctx, "ERROR", "err", err)
			// 	return nil
			// }

			return resp
		}

		return h
	}

	return m
}
