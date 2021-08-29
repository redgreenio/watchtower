package main

import (
  _ "github.com/spf13/cobra"
  _ "github.com/spf13/viper"
  "watchtower/cmd"
)

func main() {
  cmd.Execute()
}
