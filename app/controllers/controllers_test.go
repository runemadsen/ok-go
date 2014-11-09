package test_test

import (
  "github.com/go-martini/martini"
  config "golang-rails-template/config"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "net/http"
  "net/http/httptest"
  "testing"
)

var (
  m *martini.ClassicMartini
  response *httptest.ResponseRecorder
)

func TestTest(t *testing.T) {
  RegisterFailHandler(Fail)
  RunSpecs(t, "Test Suite")
}

var _ = BeforeSuite(func() {
  m = config.CreateApplication()
})

func Request(method string, route string) {
  request, _ := http.NewRequest(method, route, nil)
  response = httptest.NewRecorder()
  m.ServeHTTP(response, request)
}