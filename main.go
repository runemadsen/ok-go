package main

import (
  "github.com/runemadsen/ok-go/config"
  "github.com/runemadsen/ok-go/routes"
  "os"
)

func main() {
  app := config.NewApp(".")
  routes.Setup(app)
  app.Negroni.Run(":" + os.Getenv("PORT"))
}