package controllers

import (
  "github.com/martini-contrib/render"
)

func AboutIndex(r render.Render) {
  r.HTML(200, "about/index", "")
}