package cmd

import (
  "errors"
  "fmt"
  "github.com/spf13/cobra"
  "watchtower/download"
  "watchtower/parser"
)

var verifyCmd = &cobra.Command{
  Use: "verify",
  Args: func(cmd *cobra.Command, args []string) error {
    if len(args) < 1 {
      return errors.New("missing app ID")
    }
    return nil
  },
  Run: func(cmd *cobra.Command, args []string) {
    appId := args[0]
    content := download.Download(appId)
    appListing := parser.Parse(content)
    fmt.Println(appListing.Name)
  },
}
