package parser

import "time"

type PlayStoreAppListing struct {
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
