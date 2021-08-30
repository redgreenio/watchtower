package storage

import (
  "gorm.io/gorm"
  "watchtower/parser"
)

type DefaultReleasesRepository struct {
  db *gorm.DB
}

func (r DefaultReleasesRepository) Insert(release parser.Release) bool {
  r.db.Create(&release)
  return true
}

func (r DefaultReleasesRepository) List(appId string) []parser.Release {
  var releases []parser.Release
  r.db.Where("app_id = ?", appId).Find(&releases)
  return releases
}
