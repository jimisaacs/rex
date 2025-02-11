package rex

import (
	"time"

	"github.com/ije/rex/session"
)

var defaultAPIHanlder = &APIHandler{}
var defaultSessionPool = session.NewMemorySessionPool(time.Hour / 2)
var defaultSIDStore = session.NewCookieSIDStore("")

// Default returns the default REST
func Default() *APIHandler {
	return defaultAPIHanlder
}

// Prefix adds prefix for each api path, like "v2"
func Prefix(prefix string) *APIHandler {
	return defaultAPIHanlder.Prefix(prefix)
}

// Use appends middlewares to current APIS middleware stack.
func Use(middlewares ...Handle) {
	defaultAPIHanlder.Use(middlewares...)
}

// Query adds a query api
func Query(endpoint string, handles ...Handle) {
	defaultAPIHanlder.Query(endpoint, handles...)
}

// Mutation adds a mutation api
func Mutation(endpoint string, handles ...Handle) {
	defaultAPIHanlder.Mutation(endpoint, handles...)
}
