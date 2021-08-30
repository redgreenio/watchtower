package storage

import (
  "gorm.io/gorm"
  "watchtower/parser"
)

type DefaultReleasesRepository struct {
  db *gorm.DB
}

func Insert(repository DefaultReleasesRepository, release parser.Release) bool {
  repository.db.Create(&release)
  return true
}

func List(repository DefaultReleasesRepository, appId string) []parser.Release {
  var releases []parser.Release
  repository.db.Where("app_id = ?", appId).Find(&releases)
  return releases
}
