package config

import(
  "html/template"
  "encoding/json"
  "path/filepath"
  "io/ioutil"
  "os"
  "fmt"
  "net/http"
  "strings"
)

// The AssetHeaders struct is used to set Cache-Control headers to all GET and HEAD
// requests to /assets in production. Because these assets have digested names, we
// can set the cache time really high, and use this app as origin for a CDN.
type AssetHeaders struct {
}

func NewAssetHeaders() *AssetHeaders {
  return &AssetHeaders{}
}

func (s *AssetHeaders) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
  
  // Ignore if not production
  if os.Getenv("GO_ENV") != "production" {
    next(rw, r)
    return
  }

  // Ignore all but GET and HEAD
  if r.Method != "GET" && r.Method != "HEAD" {
    next(rw, r)
    return
  }

  // Ignore everything not in /assets
  path := r.URL.Path
  if !strings.HasPrefix(path, "/assets") {
    next(rw, r)
    return
  }

  // Set asset caching to one year 
  rw.Header().Set("Cache-Control", "public, max-age=31536000")
  next(rw, r)
  
}

// These helper functions are passed to the Go templates.
// You can add more easily by adding more functions to the FuncMap.
// The asset_path function reads from the manifest.json file to return
// the path to the digested assets in production or non-digested assets in
// development.
func AssetHelpers(root string) template.FuncMap {

  // Return digested asset paths in production
  if os.Getenv("GO_ENV") == "production" {

    var manifest map[string]interface{}
    file, e := ioutil.ReadFile(filepath.Join(root, "public/assets/manifest.json"))

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