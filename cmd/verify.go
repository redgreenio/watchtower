package cmd

import (
  "errors"
  "fmt"
  "github.com/spf13/cobra"
  _ "time"
  "watchtower/download"
  "watchtower/parser"
)

var verifyCmd = &cobra.Command{
  Use: "verify",
  Args: func(cmd *cobra.Command, args []string) error {
    if len(args) < 1 {
      return errors.New("missing argument 'appId'")
    }
    return nil
  },
  Run: func(cmd *cobra.Command, args []string) {
    appId := args[0]
    content := download.Download(appId)
    appListing := parser.Parse(content)
    printListing(appListing)
  },
}

func printListing(listing *parser.PlayStoreAppListing) {
  fmt.Println("App ID           : " + listing.AppId)
  fmt.Println("App name         : " + listing.Name)
  fmt.Println("Released on      : " + listing.ReleasedOn.Format(parser.DateLayout))
  fmt.Println("Size             : " + listing.Size)
  fmt.Println("Installs         : " + listing.Installs)
  fmt.Println("Version          : " + listing.Version)
  fmt.Println("Requires Android : " + listing.RequiresAndroid)
  fmt.Println("Content rating   : " + listing.ContentRating)
  fmt.Println("Offered by       : " + listing.OfferedBy)
}
