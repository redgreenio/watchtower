package parser

import (
  "github.com/PuerkitoBio/goquery"
  _ "github.com/PuerkitoBio/goquery"
  "strings"
)

func Parse(content string) *playStoreAppListing {
  return &playStoreAppListing{
    name: getName(content),
  }
}

func getName(content string) string {
  reader, _ := goquery.NewDocumentFromReader(strings.NewReader(content))
  selection := reader.Find(".AHFaub").Find("span")
  return strings.TrimSpace(selection.Text())
}
