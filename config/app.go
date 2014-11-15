package config

import (
  "github.com/codegangsta/negroni"
  "github.com/gorilla/mux"
  "github.com/unrolled/render"
  "html/template"
  "path/filepath"
)

type App struct {
  Negroni *negroni.Negroni
  Router *mux.Router
  Render *render.Render
}

func NewApp(root string) *App {

  negroni := negroni.Classic()
  router := mux.NewRouter()
  render := render.New(render.Options{
    Directory: filepath.Join(root, "templates"),
    Layout: "layouts/layout",
    Extensions: []string{".html"},
    Funcs: []template.FuncMap{AssetHelpers(root)},
  })

  negroni.UseHandler(router)

  return &App{negroni, router, render}
}