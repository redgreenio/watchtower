package app

type AppsRepository interface {
  Exists(appId string) bool
  Insert(app App) bool
  List() []App
}
