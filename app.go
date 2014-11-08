package main

import (
  "html/template"
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
  config "golang-rails-template/config"
)

func main() {
 
  m := martini.Classic()

  // Load manifest.json to asset helpers
  manifest := config.AssetManifest()
  AssetHelpers := template.FuncMap{
    "asset_path": func(asset string) string { return "/assets/" + manifest[asset].(string) },
  }

  // Add contrib renderer. See options in docs.
  // https://github.com/martini-contrib/render
  m.Use(render.Renderer(render.Options{
    Directory: "app/views",
    Layout: "layouts/layout",
    Extensions: []string{".html"},
    Funcs: []template.FuncMap{AssetHelpers},
  }))

  // Call the initializers
  config.Initialize(m)

  // Call the routes config
  config.Routes(m)

  m.Run()
}