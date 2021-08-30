package cmd

import (
  "fmt"
  "github.com/fatih/color"
  "github.com/spf13/cobra"
  "watchtower/app"
  "watchtower/database"
  "watchtower/download"
  "watchtower/parser"
  "watchtower/release"
)

var red = color.New(color.FgRed)

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
      content, err := download.Download(app.AppId, app.Country)
      if err != nil {
        println(fmt.Sprintf("[%d/%d] Download failed for '%s'", index+1, appsCount, app.AppId))
        continue
      }
      release := parser.Parse(content)
      if release.AppId == "" {
        _, _ = red.Printf("Unable to parse HTML for '%s'", app.AppId)
        println()
        continue
      }
      releasesRepository.Insert(release)
    }
  },
}
