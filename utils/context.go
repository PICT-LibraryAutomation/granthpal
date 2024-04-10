package utils

type contextKey struct {
	name string
}

var RemoteKey = &contextKey{name: "Database"}
var TSDBKey = &contextKey{name: "TSDB"}
var SearchKey = &contextKey{name: "Search"}
var SessionManagerKey = &contextKey{name: "SessionManager"}
var AuthKey = &contextKey{name: "Auth"}
