package app

import "gorm.io/gorm"

type App struct {
  AppId   string
  Country string
  gorm.Model
}
