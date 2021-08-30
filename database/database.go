package database

import (
  "gorm.io/driver/sqlite"
  "gorm.io/gorm"
  "watchtower/app"
  "watchtower/parser"
)

const DbPath = "xwatchtower.db"

func InitDb(dsn string) *gorm.DB {
  database, _ := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
  _ = database.AutoMigrate(parser.Release{})
  _ = database.AutoMigrate(app.App{})
  return database
}
