package routes

// This file is used to setup the general route tests. It created a new application
// with a DB connection, and resets the DB before every test runs. It also loads
// the contents of the .env.test file into the OS environment.

import (
  "github.com/runemadsen/ok-go/config"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "github.com/joho/godotenv"
  "net/http"
  "net/http/httptest"
  "net/url"
  "strings"
  "testing"
  "io"
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
  
  // Load test ENV variables
  godotenv.Load("../.env.test")

  // Create a new app with routes
  app = config.NewApp("../.")
  Setup(app)
})

var _ = BeforeEach(func() {
  // Reset the DB before every test
  app.DB.Exec("DELETE FROM posts;")
})

// Function to make a HTTP request and serve the response from
// the application.
func Request(method string, route string, body io.Reader) {
  request, _ := http.NewRequest(method, route, body)
  if(method == "POST") {
    request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
  }
  response = httptest.NewRecorder()
  app.Negroni.ServeHTTP(response, request)
}

func URLEncode(vals map[string]string) *strings.Reader {
  data := url.Values{}
  for key, value := range vals {
    data.Set(key, value)
  }
  return strings.NewReader(data.Encode())
}