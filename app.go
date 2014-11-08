package main

import (
  "github.com/go-martini/martini"
  config "golang-rails-template/config"
)

func main() {
  m := martini.Classic()
  config.Routes(m)
  m.Run()
}