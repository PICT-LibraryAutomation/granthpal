package remote

import (
	"context"
	"net/http"

	"github.com/PICT-LibraryAutomation/granthpal/utils"
	"gorm.io/gorm"
)

func RemoteMiddleware(db *gorm.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), utils.RemoteKey, db)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
