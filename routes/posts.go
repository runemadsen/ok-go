package routes

import(
  "net/http"
  "github.com/runemadsen/ok-go/models"
  "fmt"
)

func PostsIndex(w http.ResponseWriter, req *http.Request) {
  posts := []models.Post{}
  App.DB.Find(&posts)
  App.Render.HTML(w, 200, "posts/index", posts)
}

func PostsNew(w http.ResponseWriter, req *http.Request) {
  App.Render.HTML(w, 200, "posts/new", "")
}

func PostsCreate(w http.ResponseWriter, req *http.Request) {
  
  // Create a new empty Post struct
  post := new(models.Post)

  // Parse form values
  err1 := req.ParseForm()
  if err1 != nil {
    fmt.Println("Cannot parse form")
  }

  // decode form values into the struct
  err2 := decoder.Decode(post, req.PostForm)
  if err2 != nil {
    fmt.Println("Cannot decode to struct")
  }

  App.DB.Create(post)
  http.Redirect(w, req, "/posts", 302)
}