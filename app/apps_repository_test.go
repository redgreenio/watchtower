package app

import (
  "github.com/stretchr/testify/assert"
  "gorm.io/driver/sqlite"
  "gorm.io/gorm"
  "testing"
  "watchtower/parser"
)

func TestExistsReturnsFalseForEmptyTable(t *testing.T) {
  // given
  repository := testRepository()

  // when & then
  assert.False(t, repository.Exists("com.netflix.ninja"))
}

func TestInsertNewAppId(t *testing.T) {
  // given
  repository := testRepository()

  // when
  inserted := repository.Insert(App{AppId: "com.netflix.ninja", Country: nil})

  // then
  assert.True(t, inserted)
  assert.True(t, repository.Exists("com.netflix.ninja"))
}

func testRepository() AppsRepository {
  return DefaultAppsRepository{db: inMemoryDb()}
}

func inMemoryDb() *gorm.DB {
  database, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
  _ = database.AutoMigrate(parser.Release{})
  _ = database.AutoMigrate(App{})
  return database
}
