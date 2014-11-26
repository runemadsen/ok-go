package config

import(
  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
  "os"
  "fmt"
)

func NewDB() *gorm.DB {
  db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
  if err != nil {
    fmt.Printf("DB connection error: %v\n", err)
  }
  return &db
}