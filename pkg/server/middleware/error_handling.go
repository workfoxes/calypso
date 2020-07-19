package middleware

import (
	"encoding/json"
	"github.com/workfoxes/gobase/pkg/config"
	error2 "github.com/workfoxes/gobase/pkg/server/error"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"runtime/debug"
)

func ErrorHandling(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil && rvr != http.ErrAbortHandler {

				logEntry := middleware.GetLogEntry(r)
				if logEntry != nil {
					logEntry.Panic(rvr, debug.Stack())
				} else {
					middleware.PrintPrettyStack(rvr)
				}
				var err error2.BaseError
				switch x := rvr.(type) {
				case string:
					log.Print("string")
					err = error2.New(x)
				case error2.BaseError:
					err = x
				case error:
					log.Print("error")
					err = error2.ParseError(x)
				default:
					log.Print("default")
					err = error2.New("Unknown panic")
				}
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					ctx := r.Context()
					config := ctx.Value("AccountCtx").(*config.Context)
					config.DB.DB.Rollback()
					requestId := middleware.GetReqID(ctx)
					err.SetRequestId(requestId)
					_ = json.NewEncoder(w).Encode(err.Json())
				}
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
