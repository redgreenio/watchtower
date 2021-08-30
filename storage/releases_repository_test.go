package storage

import (
  "github.com/stretchr/testify/assert"
  "gorm.io/driver/sqlite"
  "gorm.io/gorm"
  "testing"
  "time"
  "watchtower/parser"
)

func TestInsertIntoEmptyTable(t *testing.T) {
  // given
  database, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
  _ = database.AutoMigrate(parser.Release{})

  repository := DefaultReleasesRepository{db: database}
  assert.Empty(t, repository.List("com.netflix.ninja"))

  // when
  release := parser.Release{
    Name:            "Netflix",
    AppId:           "com.netflix.ninja",
    ReleasedOn:      time.Now(),
    Size:            "29M",
    Installs:        "50,000,000+",
    RequiresAndroid: "Varies with device",
    ContentRating:   "Varies with device",
    OfferedBy:       "Netflix, Inc.",
  }
  inserted := repository.Insert(release)

  // then
  assert.True(t, inserted)

  releases := repository.List("com.netflix.ninja")
  assert.Len(t, releases, 1)
  assert.Equal(t, releases[0].AppId, release.AppId)
}
