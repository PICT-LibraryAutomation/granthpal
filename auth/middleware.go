package auth

import (
	"context"
	"net/http"

	"github.com/PICT-LibraryAutomation/granthpal/sessions"
	"github.com/PICT-LibraryAutomation/granthpal/utils"
)

func AuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session, err := r.Cookie(COOKIE_NAME)
			if err != nil || session == nil {
				ctx := context.WithValue(r.Context(), utils.AuthKey, nil)
				r = r.WithContext(ctx)
				next.ServeHTTP(w, r)
				return
			}

			sm := r.Context().Value(utils.SessionManagerKey).(*sessions.SessionManager)
			s, err := sm.GetSession(session.Value)
			if err != nil {
				http.SetCookie(w, &http.Cookie{
					Name:     COOKIE_NAME,
					Value:    session.Value,
					Path:     "/",
					HttpOnly: true,
					MaxAge:   -1,
				})

				ctx := context.WithValue(r.Context(), utils.AuthKey, nil)
				r = r.WithContext(ctx)
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), utils.AuthKey, &AuthData{
				SessionID: session.Value,
				UserID:    s.UserID,
			})
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
