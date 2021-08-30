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
  if !r.Exists(app.AppId) {
    r.db.Create(&app)
    return true
  }
  return false
}

func (r DefaultAppsRepository) List() []App {
  var apps []App
  r.db.Find(&apps)
  return apps
}
