package routes

import(
  "net/http"
)

func AboutIndex(w http.ResponseWriter, req *http.Request) {
  App.Render.HTML(w, 200, "about/index", "")
}