package parser

import (
  "github.com/PuerkitoBio/goquery"
  _ "github.com/PuerkitoBio/goquery"
  "net/url"
  "strings"
)

func Parse(content string) *PlayStoreAppListing {
  document, _ := goquery.NewDocumentFromReader(strings.NewReader(content))

  return &PlayStoreAppListing{
    Name:  getName(document),
    AppId: getAppId(document),
  }
}

func getName(document *goquery.Document) string {
  span := document.Find(".AHFaub").Find("span")
  return strings.TrimSpace(span.Text())
}

func getAppId(document *goquery.Document) string {
  link := document.Find("[rel=canonical]")
  playStoreUrl := getPlayStoreUrl(link)
  parsedUrl, _ := url.Parse(playStoreUrl)
  return parsedUrl.Query().Get("id")
}

func getPlayStoreUrl(linkElement *goquery.Selection) string {
  playStoreUrl, _ := linkElement.Attr("href")
  return playStoreUrl
}
