package utils

type contextKey struct {
	name string
}

var RemoteKey = &contextKey{name: "Database"}
var SearchKey = &contextKey{name: "Search"}
var SessionManagerKey = &contextKey{name: "SessionManager"}
var AuthKey = &contextKey{name: "Auth"}
