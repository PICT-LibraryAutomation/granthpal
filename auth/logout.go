package auth

import (
	"net/http"

	"github.com/PICT-LibraryAutomation/granthpal/sessions"
	"github.com/PICT-LibraryAutomation/granthpal/utils"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie(COOKIE_NAME)
	if err == http.ErrNoCookie {
		http.Error(w, "No user logged in", http.StatusUnauthorized)
		return
	}

	sm := r.Context().Value(utils.SessionManagerKey).(*sessions.SessionManager)
	err = sm.DeleteSession(session.Value)
	if err != nil {
		http.Error(w, "Couldn't log out", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     COOKIE_NAME,
		Value:    session.Value,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	})

	w.Write([]byte("Logged Out"))
}
