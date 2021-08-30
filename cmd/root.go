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
  // add <appId>
  addCmd.SetUsageTemplate("hint: Maybe you wanted to say 'watchtower add <appId>'?")
  rootCmd.AddCommand(addCmd)

  // list
  rootCmd.AddCommand(listCmd)

  // verify <appId>
  verifyCmd.SetUsageTemplate("hint: Maybe you wanted to say 'watchtower verify <appId>'?")
  rootCmd.AddCommand(verifyCmd)

  // observe
  rootCmd.AddCommand(observeCmd)
}
