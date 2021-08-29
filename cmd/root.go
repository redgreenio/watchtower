package cmd

import (
  "github.com/spf13/cobra"
  "os"
)

var rootCmd = &cobra.Command{
  Use: "watchtower",
  Run: func(cmd *cobra.Command, args []string) {
    // Nothing in here yet!
  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    os.Exit(1)
  }
}

func init() {
  rootCmd.AddCommand(verifyCmd)
}
