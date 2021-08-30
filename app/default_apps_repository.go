package app

import "gorm.io/gorm"

type DefaultAppsRepository struct {
  db *gorm.DB
}

func (r DefaultAppsRepository) Exists(appId string) bool {
  return false
}