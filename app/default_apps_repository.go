package app

import "gorm.io/gorm"

type DefaultAppsRepository struct {
  db *gorm.DB
}

func (r DefaultAppsRepository) Exists(appId string) bool {
  var app *App
  r.db.Where("app_id = ?", appId).Find(&app)
  return app.AppId == appId
}

func (r DefaultAppsRepository) Insert(app App) bool {
  r.db.Create(&app)
  return true
}
