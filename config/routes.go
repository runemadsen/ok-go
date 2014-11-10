package config

import (
  "github.com/go-martini/martini"
  controllers "golang-rails-template/app/controllers"
)

func Routes(m *martini.ClassicMartini) {
  
  m.Get("/", controllers.HomeIndex)
  
  m.Get("/about", controllers.AboutIndex)

}