package rex

import (
	"time"

	"github.com/ije/rex/session"
)

var defaultREST = New("/")
var defaultSessionPool = session.NewMemorySessionPool(time.Hour / 2)
var defaultSIDStore = &session.CookieSIDStore{}

// Default returns the default REST
func Default() *REST {
	return defaultREST
}

// Group creates a nested REST
func Group(path string, nest func(*REST)) *REST {
	return defaultREST.Group(path, nest)
}

// Use appends middleware to the REST middleware stack.
func Use(middlewares ...Handle) {
	defaultREST.Use(middlewares...)
}

// Fallback handles the out-routed requests.
func Fallback(handle Handle) {
	defaultREST.Fallback(handle)
}

// Options is a shortcut for rest.Handle("OPTIONS", path, handles)
func Options(path string, handles ...Handle) {
	defaultREST.Options(path, handles...)
}

// Head is a shortcut for rest.Handle("HEAD", path, handles)
func Head(path string, handles ...Handle) {
	defaultREST.Head(path, handles...)
}

// Get is a shortcut for rest.Handle("GET", path, handles)
func Get(path string, handles ...Handle) {
	defaultREST.Get(path, handles...)
}

// Post is a shortcut for rest.Handle("POST", path, handles)
func Post(path string, handles ...Handle) {
	defaultREST.Post(path, handles...)
}

// Put is a shortcut for rest.Handle("PUT", path, handles)
func Put(path string, handles ...Handle) {
	defaultREST.Put(path, handles...)
}

// Patch is a shortcut for rest.Handle("PATCH", path, handles)
func Patch(path string, handles ...Handle) {
	defaultREST.Patch(path, handles...)
}

// Delete is a shortcut for rest.Handle("DELETE", path, handles)
func Delete(path string, handles ...Handle) {
	defaultREST.Delete(path, handles...)
}

// Trace is a shortcut for rest.Handle("TRACE", path, handles)
func Trace(path string, handles ...Handle) {
	defaultREST.Trace(path, handles...)
}

// Static handles static files requests.
func Static(path string, root string, fallbackPath ...string) {
	defaultREST.Static(path, root, fallbackPath...)
}
