package models

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("Post", func() {

  It("saves a post", func() {

    // Make sure we have 0 posts
    posts := []Post{}
    app.DB.Find(&posts)
    Expect(len(posts)).To(BeZero())
    
    // Create a new post
    app.DB.Create(Post{Title:"My Post", Body:"My Body"})
    
    // Make sure we have 1 post
    app.DB.Find(&posts)
    Expect(len(posts)).To(Equal(1))
      
    // Make sure that post is the one we created
    post := Post{}
    app.DB.First(&post)
    Expect(post.Title).To(Equal("My Post"))
    Expect(post.Body).To(Equal("My Body"))
  })

})