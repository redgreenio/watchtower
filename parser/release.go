package parser

import (
  "gorm.io/gorm"
  "time"
)

const ReleasedOnDateLayout string = "January 2, 2006"

type Release struct {
  Name            string
  AppId           string
  ReleasedOn      time.Time
  Size            string
  Installs        string
  Version         string
  RequiresAndroid string
  ContentRating   string
  OfferedBy       string
  gorm.Model
}
