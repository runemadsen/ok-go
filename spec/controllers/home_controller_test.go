package controllers_test

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "fmt"
)

var _ = Describe("Home", func() {

  It("renders the home page", func() {    
    Request("GET", "/")
    fmt.Println(response.Body)
    Expect(response.Code).To(Equal(200))
  })

})