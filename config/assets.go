package config

import (
  "encoding/json"
  "io/ioutil"
  "fmt"
)

func AssetManifest() map[string]interface{} {

  file, e := ioutil.ReadFile("public/assets/manifest.json")
  if e != nil {
    fmt.Printf("File error: %v\n", e)
  }

  var digest map[string]interface{}
  err := json.Unmarshal(file, &digest)
  if err != nil {
    fmt.Printf("JSON unmarshal error: %v\n", err)
  }

  return digest
}