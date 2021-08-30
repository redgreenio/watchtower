package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
  "watchtower/app"
  "watchtower/database"
)

var listCmd = &cobra.Command{
  Use: "list",
  Run: func(cmd *cobra.Command, args []string) {
    database := database.InitDb(database.DbPath)
    repository := app.DefaultAppsRepository{Database: database}
    apps := repository.List()

    for _, app := range apps {
      if app.Country != "" {
        println(fmt.Sprintf("%s (%s)", app.AppId, app.Country))
      } else {
        println(app.AppId)
      }
    }

    appsCount := len(apps)
    if appsCount == 0 {
      println("Uh oh... such empty.")
      return
    }

    println("~~")
    if appsCount == 1 {
      println("1 app found.")
    } else {
      println(fmt.Sprintf("%d apps found.", appsCount))
    }
  },
}
