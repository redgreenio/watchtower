package release

import (
  "gorm.io/gorm"
  "watchtower/parser"
)

type DefaultReleasesRepository struct {
  Database *gorm.DB
}

func (r DefaultReleasesRepository) Insert(release parser.Release) bool {
  var existingRelease *parser.Release
  r.Database.Where("app_id = ? AND version = ? AND released_on = ?", release.AppId, release.Version, release.ReleasedOn).Find(&existingRelease)

  if existingRelease.AppId == "" {
    r.Database.Create(&release)
    return true
  } else {
    return false
  }
}

func (r DefaultReleasesRepository) List(appId string) []parser.Release {
  var releases []parser.Release
  r.Database.Where("app_id = ?", appId).Find(&releases)
  return releases
}
