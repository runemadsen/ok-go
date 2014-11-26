package routes

import (
  "github.com/runemadsen/ok-go/models"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("Posts", func() {

  Describe("#PostsIndex", func() {

    It("renders a list of posts", func() {    
      
      app.DB.Create(models.Post{Title:"First Post", Body:"First Body"})
      app.DB.Create(models.Post{Title:"Second Post", Body:"Second Body"})

      Request("GET", "/posts", nil)
      Expect(response.Code).To(Equal(200))
      Expect(response.Body).To(ContainSubstring("<h1>Posts</h1>"))
      Expect(response.Body).To(ContainSubstring("<h2>First Post</h2>"))
      Expect(response.Body).To(ContainSubstring("<h2>Second Post</h2>"))
    })
  
  })

  Describe("#PostsCreate", func() {

    It("creates a new post", func() {    
      
      posts := []models.Post{}
      app.DB.Find(&posts)
      Expect(len(posts)).To(BeZero())

      Request("POST", "/posts", URLEncode(map[string]string{
        "title" : "My Post",
        "body"  : "My Body",
      }))
      
      Expect(response.Code).To(Equal(302))
      app.DB.Find(&posts)
      Expect(len(posts)).To(Equal(1))
      
      post := models.Post{}
      app.DB.First(&post)
      Expect(post.Title).To(Equal("My Post"))
      Expect(post.Body).To(Equal("My Body"))
    })
  
  })

})