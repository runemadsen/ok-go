package main

import (
  "github.com/runemadsen/ok-go/config"
  "github.com/runemadsen/ok-go/routes"
)

func main() {
  app := config.NewApp(".")
  routes.Setup(app)
  app.Negroni.Run(":3000")
}