package config

import(
  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
  "github.com/runemadsen/ok-go/models"
  "os"
  "fmt"
)

func NewDB() *gorm.DB {
  db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
  if err != nil {
    fmt.Printf("DB connection error: %v\n", err)
  }
  runMigrations(&db)
  return &db
}

func runMigrations(db *gorm.DB) {
  db.CreateTable(&models.Post{})
}