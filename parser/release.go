package parser

import "time"

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
}
