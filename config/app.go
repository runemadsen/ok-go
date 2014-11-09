package config

import (
  "html/template"
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
)

func CreateApplication() *martini.ClassicMartini {

  m := martini.Classic()

  // Add contrib renderer. See options in docs.
  // https://github.com/martini-contrib/render
  m.Use(render.Renderer(render.Options{
    Directory: "app/views",
    Layout: "layouts/layout",
    Extensions: []string{".html"},
    Funcs: []template.FuncMap{AssetHelpers()},
  }))

  // Call the initializers
  Initialize(m)

  // Call the routes
  Routes(m)

  return m

}