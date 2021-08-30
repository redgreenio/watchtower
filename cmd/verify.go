package cmd

import (
  "errors"
  "github.com/fatih/color"
  _ "github.com/fatih/color"
  "github.com/spf13/cobra"
  _ "time"
  "watchtower/download"
  "watchtower/parser"
)

const appId = "App ID"
const appName = "App name"
const releasedOn = "Updated on"
const size = "Size"
const installs = "Installs"
const version = "Version"
const requiresAndroid = "Requires Android"
const contentRating = "Content rating"
const offeredBy = "Offered by"

const messageListingFound = "Listing found ✓"
const formatTitleValue = "%-18s"

var green = color.New(color.Bold, color.FgGreen)
var bold = color.New(color.Bold)

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

func printListing(listing *parser.Release) {
  printListingFound()
  printFormattedLine(appId, listing.AppId)
  printFormattedLine(appName, listing.Name)
  printFormattedLine(releasedOn, listing.ReleasedOn.Format(parser.DateLayout))
  printFormattedLine(size, listing.Size)
  printFormattedLine(installs, listing.Installs)
  printFormattedLine(version, listing.Version)
  printFormattedLine(requiresAndroid, listing.RequiresAndroid)
  printFormattedLine(contentRating, listing.ContentRating)
  printFormattedLine(offeredBy, listing.OfferedBy)
}

func printListingFound() {
  _, _ = green.Println(messageListingFound)
}

func printFormattedLine(title, value string) {
  _, _ = bold.Printf(formatTitleValue, title)
  println(value)
}
