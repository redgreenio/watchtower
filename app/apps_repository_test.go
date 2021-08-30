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
  inserted := repository.Insert(App{AppId: "com.netflix.ninja", Country: ""})

  // then
  assert.True(t, inserted)
  assert.True(t, repository.Exists("com.netflix.ninja"))
}

func TestDoNotInsertExistingAppId(t *testing.T) {
  // given
  repository := testRepository()
  repository.Insert(App{AppId: "com.netflix.ninja", Country: ""})

  // when
  inserted := repository.Insert(App{AppId: "com.netflix.ninja", Country: ""})

  // then
  assert.False(t, inserted)
  assert.True(t, repository.Exists("com.netflix.ninja"))
}

func TestNoEntriesInAnEmptyTable(t *testing.T) {
  // given
  repository := testRepository()

  // when
  apps := repository.List()

  // then
  assert.Len(t, apps, 0)
}

func TestListAllApps(t *testing.T) {
  // given
  repository := testRepository()
  repository.Insert(App{AppId: "com.netflix.ninja"})
  repository.Insert(App{AppId: "io.redgreen.watchtower", Country: "IN"})

  // when
  apps := repository.List()

  // then
  assert.Len(t, apps, 2)
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
