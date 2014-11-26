package models

import (
  "github.com/runemadsen/ok-go/config"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "github.com/joho/godotenv"
  "testing"
)

var (
  app *config.App
)

func TestTest(t *testing.T) {
  RegisterFailHandler(Fail)
  RunSpecs(t, "Test Suite")
}

var _ = BeforeSuite(func() {
  
  // Load test ENV variables
  godotenv.Load("../.env.test")

  // Create a new app
  app = config.NewApp("../.")
})

var _ = BeforeEach(func() {
  // Reset the DB before every test
  app.DB.Exec("DELETE FROM posts;")
})