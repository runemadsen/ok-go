package models

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("Post", func() {

  It("saves a post", func() {

    posts := []Post{}
    app.DB.Find(&posts)
    Expect(len(posts)).To(BeZero())
    
    app.DB.Create(Post{Title:"My Post", Body:"My Body"})
    
    app.DB.Find(&posts)
    Expect(len(posts)).To(Equal(1))
      
    post := Post{}
    app.DB.First(&post)
    Expect(post.Title).To(Equal("My Post"))
    Expect(post.Body).To(Equal("My Body"))
  })

})