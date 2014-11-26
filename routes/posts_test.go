package routes

import (
  "github.com/runemadsen/ok-go/models"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("Posts", func() {

  Describe("#PostsIndex", func() {

    It("renders as list of posts", func() {    
      
      app.DB.Create(models.Post{Title:"First Post", Body:"First Body"})
      app.DB.Create(models.Post{Title:"Second Post", Body:"Second Body"})

      Request("GET", "/posts")
      Expect(response.Code).To(Equal(200))
      Expect(response.Body).To(ContainSubstring("<h1>Posts</h1>"))
      Expect(response.Body).To(ContainSubstring("<h2>First Post</h2>"))
      Expect(response.Body).To(ContainSubstring("<h2>Second Post</h2>"))
    })
  
  })

})