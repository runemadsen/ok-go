package config

import (
  "html/template"
  "encoding/json"
  "io/ioutil"
  "os"
  "fmt"
  "path/filepath"
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
)

func CreateApplication(basePath string) *martini.ClassicMartini {

  m := martini.Classic()

  m.Use(render.Renderer(render.Options{
    Directory: filepath.Join(basePath, "app/views"),
    Layout: "layouts/layout",
    Extensions: []string{".html"},
    Funcs: []template.FuncMap{AssetHelpers(basePath)},
  }))

  Initialize(m)
  Routes(m)

  return m
}

func AssetHelpers(basePath string) template.FuncMap {

  // Return digested asset paths in production
  if os.Getenv("MARTINI_ENV") == "production" {

    var manifest map[string]interface{}
    file, e := ioutil.ReadFile(filepath.Join(basePath, "public/assets/manifest.json"))

    if e != nil {
      fmt.Printf("File error: %v\n", e)
    }
  
    err := json.Unmarshal(file, &manifest)
    if err != nil {
      fmt.Printf("JSON unmarshal error: %v\n", err)
    }
  
    return template.FuncMap{
      "asset_path": func(asset string) string { 
        return "/assets/" + manifest[asset].(string)
      },
    }

  // Return non-digested asset paths in other envs
  } else {
    return template.FuncMap{
      "asset_path": func(asset string) string { 
        return "/assets/" + asset
      },
    }
  }  
}