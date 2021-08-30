package app

type AppsRepository interface {
  Exists(appId string) bool
}
