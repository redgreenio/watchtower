package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
  "watchtower/app"
  "watchtower/database"
  "watchtower/download"
  "watchtower/parser"
  "watchtower/release"
)

var observeCmd = &cobra.Command{
  Use: "observe",
  Run: func(cmd *cobra.Command, args []string) {
    database := database.InitDb(database.DbPath)
    appsRepository := app.DefaultAppsRepository{Database: database}
    releasesRepository := release.DefaultReleasesRepository{Database: database}

    apps := appsRepository.List()
    appsCount := len(apps)
    for index, app := range apps {
      println(fmt.Sprintf("[%d/%d] Downloading release info for '%s'", index+1, appsCount, app.AppId))
      content := download.Download(app.AppId)
      release := parser.Parse(content)
      releasesRepository.Insert(release)
    }
  },
}
