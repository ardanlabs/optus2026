package mid

import (
	"context"
	"errors"
	"net/http"
	"path"

	"github.com/ardanlabs/kronk/cmd/server/foundation/logger"
	"github.com/ardanlabs/service/app/sdk/errs"
	"github.com/ardanlabs/service/foundation/web"
)

func checkIsError(e web.Encoder) error {
	err, hasError := e.(error)
	if hasError {
		return err
	}

	return nil
}

func Errors(log *logger.Logger) web.MidFunc {
	m := func(handler web.HandlerFunc) web.HandlerFunc {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) web.Encoder {
			resp := handler(ctx, w, r)

			err := checkIsError(resp)
			if err == nil {
				return resp
			}

			appErr, ok := errors.AsType[*errs.Error](err)
			if !ok {
				appErr = errs.Errorf(errs.Internal, "Internal Server Error")
			}

			log.Error(ctx, "handled error during request",
				"err", err,
				"source_err_file", path.Base(appErr.FileName),
				"source_err_func", path.Base(appErr.FuncName))

			if appErr.Code == errs.InternalOnlyLog {
				appErr = errs.Errorf(errs.Internal, "Internal Server Error")
			}

			// Send the error to the web package so the error can be
			// used as the response.

			return appErr
		}

		return h
	}

	return m
}
