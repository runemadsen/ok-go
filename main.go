package main

import (
  config "golang-rails-template/config"
)

func main() {
  m := config.CreateApplication()
  m.Run()
}