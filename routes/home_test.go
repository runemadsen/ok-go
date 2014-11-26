package routes

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("Home", func() {

  It("renders the home page", func() {    
    Request("GET", "/", nil)
    Expect(response.Code).To(Equal(200))
    Expect(response.Body).To(ContainSubstring("This is the home view"))
  })

})