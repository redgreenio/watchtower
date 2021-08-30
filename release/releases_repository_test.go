package release

import (
  "github.com/stretchr/testify/assert"
  "gorm.io/gorm"
  "testing"
  "time"
  "watchtower/database"
  "watchtower/parser"
)

var releasedOn, _ = time.Parse(parser.ReleasedOnDateLayout, "August 9, 2021")
var testRelease = parser.Release{
  Name:            "Netflix",
  AppId:           "com.netflix.ninja",
  ReleasedOn:      releasedOn,
  Size:            "29M",
  Installs:        "50,000,000+",
  Version:         "12.0.1",
  RequiresAndroid: "Varies with device",
  ContentRating:   "Varies with device",
  OfferedBy:       "Netflix, Inc.",
}

func TestInsertIntoEmptyTable(t *testing.T) {
  // given
  repository := testRepository()
  assert.Empty(t, repository.List("com.netflix.ninja"))

  // when
  inserted := repository.Insert(testRelease)

  // then
  assert.True(t, inserted)

  releases := repository.List("com.netflix.ninja")
  assert.Len(t, releases, 1)
  assert.Equal(t, releases[0].AppId, testRelease.AppId)
}

func TestDoNotInsertReleaseWithSameVersion(t *testing.T) {
  // given
  repository := testRepository()
  repository.Insert(testRelease)

  // when
  inserted := repository.Insert(testRelease)

  // then
  assert.False(t, inserted)
  releases := repository.List("com.netflix.ninja")
  assert.Len(t, releases, 1)
}

func TestInsertReleaseWithDifferentVersion(t *testing.T) {
  // given
  repository := testRepository()
  repository.Insert(testRelease)

  // when
  testRelease.Version = "12.0.2"
  inserted := repository.Insert(testRelease)

  // then
  assert.True(t, inserted)
  releases := repository.List("com.netflix.ninja")
  assert.Len(t, releases, 2)
}

func TestInsertReleaseWithDifferentReleasedOnDate(t *testing.T) {
  // given
  repository := testRepository()
  repository.Insert(testRelease)
  newerReleasedOn, _ := time.Parse(parser.ReleasedOnDateLayout, "August 10, 2021")

  // when
  testRelease.ReleasedOn = newerReleasedOn
  inserted := repository.Insert(testRelease)

  // then
  assert.True(t, inserted)
  releases := repository.List("com.netflix.ninja")
  assert.Len(t, releases, 2)
}

func testRepository() ReleasesRepository {
  return DefaultReleasesRepository{db: inMemoryDb()}
}

func inMemoryDb() *gorm.DB {
  return database.InitDb("file::memory:")
}
