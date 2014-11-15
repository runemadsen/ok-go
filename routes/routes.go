package routes

import(
  "github.com/runemadsen/ok-go/config"
)

var App *config.App

func Setup(app *config.App) {
  
  App = app

  // Define your routes here:

  App.Router.HandleFunc("/", HomeIndex).Methods("GET")
  App.Router.HandleFunc("/about", AboutIndex).Methods("GET")
}