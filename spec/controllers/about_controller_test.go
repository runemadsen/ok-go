package controllers_test

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("Home", func() {

  It("renders the about page", func() {    
    Request("GET", "/about")
    Expect(response.Code).To(Equal(200))
    Expect(response.Body).To(ContainSubstring("This is the about view"))
  })

})