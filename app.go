package main

import (
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
  config "golang-rails-template/config"
)

func main() {
 
  m := martini.Classic()

  // Add contrib renderer. See options in docs.
  // https://github.com/martini-contrib/render
  m.Use(render.Renderer(render.Options{
    Directory: "app/views",
    Layout: "layouts/layout",
    Extensions: []string{".html"},
  }))

  // Call the routes config
  config.Routes(m)

  m.Run()
}