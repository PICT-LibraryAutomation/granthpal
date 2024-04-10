package auth

import "time"

const HASH_SALTS = 10
const COOKIE_NAME = "session-id"
const COOKIE_DURATION = (24 * time.Hour) * 100 // 100 days duration

type AuthData struct {
	SessionID string
	UserID    string
}
