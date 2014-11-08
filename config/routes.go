package config

import (
  "github.com/go-martini/martini"
  controllers "golang-rails-template/app/controllers"
)

func Routes(m *martini.ClassicMartini) {
  
  // Define your routes here
  m.Get("/", controllers.HomeIndex)
  m.Get("/about", controllers.AboutIndex)

}