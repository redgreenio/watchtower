package main

import (
  "fmt"
  "log"
  "net/http"
)
import _ "net/http"

func main() {
  resp, err := http.Get("https://google.com")
  if err != nil {
    log.Fatalln("Uh oh! unable to fetch URL.")
  }
  fmt.Println(resp.Status)
}
