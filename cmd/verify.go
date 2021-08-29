package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
  "watchtower/download"
  "watchtower/parser"
)

var verifyCmd = &cobra.Command{
  Use: "verify",
  Run: func(cmd *cobra.Command, args []string) {
    content := download.Download("com.spotify.music")
    appListing := parser.Parse(content)
    fmt.Println(appListing.Name)
  },
}
