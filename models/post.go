package models

import(
  "time"
)

type Post struct {
  Id         int64
  Title      string  `sql:"size:255"`
  Body       string
  CreatedAt  time.Time
  UpdatedAt  time.Time
}