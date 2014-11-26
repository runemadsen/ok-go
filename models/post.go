package models

import(
  "time"
)

// This is our single struct/model that is used with Gorm for all 
// DB functions.
type Post struct {
  Id         int64
  Title      string  `sql:"size:255"`
  Body       string
  CreatedAt  time.Time
  UpdatedAt  time.Time
}