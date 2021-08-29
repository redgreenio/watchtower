package main

import (
  "fmt"
  "watchtower/download"
  "watchtower/parser"
)

func main() {
  content := download.Download("com.spotify.music")
  appListing := parser.Parse(content)
  fmt.Println(appListing.Name)
}
