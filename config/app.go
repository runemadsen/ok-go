package config

import (
  "github.com/codegangsta/negroni"
  "github.com/gorilla/mux"
  "github.com/unrolled/render"
  "github.com/jinzhu/gorm"
  "html/template"
  "path/filepath"
  "net/http"
)

type App struct {
  Negroni *negroni.Negroni
  Router *mux.Router
  Render *render.Render
  DB *gorm.DB
}

func NewApp(root string) *App {

  CheckEnv()

  ne := negroni.New()
  ro := mux.NewRouter()
  re := render.New(render.Options{
    Directory: filepath.Join(root, "templates"),
    Layout: "layouts/layout",
    Extensions: []string{".html"},
    Funcs: []template.FuncMap{AssetHelpers(root)},
  })
  db := NewDB()

  ne.Use(negroni.NewRecovery())
  ne.Use(negroni.NewLogger())
  ne.Use(NewAssetHeaders())
  ne.Use(negroni.NewStatic(http.Dir("public")))
  ne.UseHandler(ro)

  return &App{ne, ro, re, db}
}