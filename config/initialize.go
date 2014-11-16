package config

import(
  "os"
  "log"
)

// Check that we have required ENV variables to run the app
// Add names of ENV variables to env[] slice as you add them
// to your application logic.
func init() {

  env := []string{
    "DATABASE_URL",
  }

  for _, value := range env {
    if os.Getenv(value) == "" {
      log.Fatal("ENV variable not provided: " + value)
    }
  }

}