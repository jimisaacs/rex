package main

import (
	"github.com/jimisaacs/rex"
)

func main() {
	rex.Use(rex.AutoCompress())

	rex.Query("*", func(ctx *rex.Context) interface{} {
		return rex.FS("./www", "e404.html")
	})

	<-rex.Start(8080)
}
