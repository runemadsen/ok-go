package routes

import(
  "net/http"
)

func HomeIndex(w http.ResponseWriter, req *http.Request) {
  App.Render.HTML(w, 200, "home/index", nil)
}