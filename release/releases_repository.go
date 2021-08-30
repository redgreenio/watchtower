package release

import (
  "watchtower/parser"
)

type ReleasesRepository interface {
  Insert(release parser.Release) bool
  List(appId string) []parser.Release
}
