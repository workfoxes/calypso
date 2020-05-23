package middleware

import (
	"context"
	"github.com/airstrik/gobase/pkg/config"
	"github.com/go-chi/chi"
	"net/http"
)

func Autherization(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		_CSRFToken := r.Header.Get("")
		_AccountId := chi.URLParam(r, "AccountId")
		if  _CSRFToken != "" {
			// Here Comes the Account Validation and Autherization
		}
		ctx = context.WithValue(ctx, "context", config.NewContext(_AccountId))
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}