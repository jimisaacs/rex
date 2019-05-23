package rex

import (
	"net/url"

	"github.com/ije/rex/router"
)

// A URL is a *url.URL with httprouter.Params
type URL struct {
	Params router.Params
	*url.URL
}

// Param returns the value of the first Param which key matches the given name.
// If no matching Param is found, an empty string is returned.
func (url *URL) Param(name string) string {
	return url.Params[name]
}
