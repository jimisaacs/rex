package rex

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/ije/rex/router"
	"github.com/ije/rex/session"
)

var defaultREST = New()
var defaultRouter = router.New()
var defaultSessionPool = session.NewMemorySessionPool(time.Hour / 2)
var defaultSIDStore = &session.CookieSIDStore{}

func init() {
	defaultRouter.HandleOptions(func(w http.ResponseWriter, r *http.Request) {
		defaultREST.serve(w, r, nil, func(ctx *Context) {
			ctx.End(http.StatusNoContent)
		})
	})
	defaultRouter.HandlePanic(func(w http.ResponseWriter, r *http.Request, v interface{}) {
		if err, ok := v.(*contextPanicError); ok {
			defaultREST.serve(w, r, nil, func(ctx *Context) {
				ctx.Error(err.message, err.code)
			})
			return
		}

		buf := bytes.NewBuffer(nil)
		for i := 3; ; i++ {
			pc, file, line, ok := runtime.Caller(i)
			if !ok {
				break
			}
			fmt.Fprint(buf, "\t", strings.TrimSpace(runtime.FuncForPC(pc).Name()), " ", file, ":", line, "\n")
		}

		defaultREST.serve(w, r, nil, func(ctx *Context) {
			ctx.Error(fmt.Sprintf("[panic] %v\n%s", v, buf.String()), 500)
		})
	})
}

// NotFound sets a not found handle
func NotFound(handle Handle) {
	defaultRouter.NotFound(func(w http.ResponseWriter, r *http.Request) {
		defaultREST.serve(w, r, nil, handle)
	})
}

// Group creates a nested REST
func Group(prefix string, callback func(*REST)) *REST {
	return defaultREST.Group(prefix, callback)
}

// Use appends middleware to the REST middleware stack.
func Use(middlewares ...Handle) {
	defaultREST.Use(middlewares...)
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
