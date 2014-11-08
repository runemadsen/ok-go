package config

import (
  "html/template"
  "encoding/json"
  "io/ioutil"
  "os"
  "fmt"
)

func AssetHelpers() template.FuncMap {

  // Return digested asset paths in production
  if os.Getenv("MARTINI_ENV") == "production" {

    var manifest map[string]interface{}
    file, e := ioutil.ReadFile("public/assets/manifest.json")
    if e != nil {
      fmt.Printf("File error: %v\n", e)
    }
  
    err := json.Unmarshal(file, &manifest)
    if err != nil {
      fmt.Printf("JSON unmarshal error: %v\n", err)
    }
  
    return template.FuncMap{
      "asset_path": func(asset string) string { 
        if os.Getenv("MARTINI_ENV") == "production" {
          return "/assets/" + manifest[asset].(string)
        } else {
          return asset
        }
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