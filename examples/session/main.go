package main

import (
	"time"

	"github.com/ije/rex"
	"github.com/ije/rex/session"
)

const indexHTML = `
<h1>Welcome to use REX!</h1>{{if .user}}
<p>Hello, {{.user}}</p>
<p>You will logout after 15 seconds automatically...</p>
<p><a href="/logout">Logout</a></p>{{else}}
<form method="post" action="/login">
	<label>Login as:</label>
	<input name="user" type="text">
</form>{{end}}
`

func main() {
	rex.Get("/", func(ctx *rex.Context) {
		ctx.RenderHTML(indexHTML, map[string]interface{}{
			"user": ctx.Session().Get("user"),
		})
	})

	rex.Post("/login", func(ctx *rex.Context) {
		user := ctx.Form.Get("user")
		if user != "" {
			ctx.Session().Set("user", user)
		}
		ctx.Redirect("/", 301)
	})

	rex.Get(
		"/logout",
		rex.Header("Cache-Control", "no-cache, no-store, must-revalidate"),
		func(ctx *rex.Context) {
			ctx.Session().Delete("user")
			ctx.Redirect("/", 301)
		},
	)

	rex.Use(
		rex.SessionPool(session.NewMemorySessionPool(15*time.Second)),
		rex.SendError(),
	)
	rex.Start(8080)
}
