package cmd

import (
  "errors"
  "fmt"
  "github.com/spf13/cobra"
  "watchtower/app"
  "watchtower/database"
)

var addCmd = &cobra.Command{
  Use: "add",
  Args: func(cmd *cobra.Command, args []string) error {
    if len(args) < 1 {
      return errors.New("missing argument 'appId'")
    }
    return nil
  },
  Run: func(cmd *cobra.Command, args []string) {
    appId := args[0]
    database := database.InitDb(database.DbPath)
    repository := app.DefaultAppsRepository{Database: database}
    inserted := repository.Insert(app.App{AppId: appId})

    if inserted {
      println(fmt.Sprintf("'%s' added to list.", appId))
    } else {
      println(fmt.Sprintf("'%s' is already on the list.", appId))
    }
  },
}
