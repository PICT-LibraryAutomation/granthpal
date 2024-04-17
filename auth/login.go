package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/PICT-LibraryAutomation/granthpal/remote/models"
	"github.com/PICT-LibraryAutomation/granthpal/sessions"
	"github.com/PICT-LibraryAutomation/granthpal/utils"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type loginBody struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var b loginBody
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, "Couldn't parse request body", http.StatusBadRequest)
		return
	}

	_, err := r.Cookie(COOKIE_NAME)
	if err != http.ErrNoCookie {
		http.Error(w, "User already logged in", http.StatusUnauthorized)
		return
	}

	sm := r.Context().Value(utils.SessionManagerKey).(*sessions.SessionManager)
	remDB := r.Context().Value(utils.RemoteKey).(*gorm.DB)

	var u models.User
	if err := remDB.Find(&u, "id = ?", b.UserID).Error; err != nil {
		http.Error(w, "No such user", http.StatusBadRequest)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(b.Password)); err != nil {
		http.Error(w, "Incorrect password provided", http.StatusBadRequest)
		return
	}

	sessionID := uuid.NewString()
	expiry := time.Now().Add(COOKIE_DURATION)

	err = sm.NewSession(sessions.Session{
		ID:     sessionID,
		UserID: b.UserID,
		Expiry: expiry,
	})
	if err != nil {
		http.Error(w, "Couldn't create a new session", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     COOKIE_NAME,
		Value:    sessionID,
		Expires:  expiry,
		Path:     "/",
		HttpOnly: true,
	})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Authenticated"))
}
