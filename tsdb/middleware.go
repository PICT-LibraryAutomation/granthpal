package tsdb

import (
	"context"
	"net/http"

	"github.com/PICT-LibraryAutomation/granthpal/utils"
)

func TSDBMiddleware(t *TSDB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), utils.TSDBKey, t)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
