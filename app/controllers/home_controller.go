package controllers

import (
  "github.com/martini-contrib/render"
)

func HomeIndex(r render.Render) {
  r.HTML(200, "home/index", "")
}