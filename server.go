package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func main() {
	m := martini.Classic()

	//中间件
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))

	m.Get("/", func() string {
		return "Hello world!"
	})

	m.Get("/render", func(r render.Render) {
		r.HTML(200, "hello", "jeremy")
	})
	m.Run()
}