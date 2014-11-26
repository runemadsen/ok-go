package routes

import (
  "github.com/runemadsen/ok-go/config"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "github.com/joho/godotenv"
  "net/http"
  "net/http/httptest"
  "testing"
)

var (
  app *config.App
  response *httptest.ResponseRecorder
)

func TestTest(t *testing.T) {
  RegisterFailHandler(Fail)
  RunSpecs(t, "Test Suite")
}

var _ = BeforeSuite(func() {
  godotenv.Load("../.env.test")
  app = config.NewApp("../.")
  Setup(app)
})

func Request(method string, route string) {
  request, _ := http.NewRequest(method, route, nil)
  response = httptest.NewRecorder()
  app.Negroni.ServeHTTP(response, request)
}