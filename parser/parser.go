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
    Name: getName(document),
    AppId: getAppId(document),
  }
}

func getName(document *goquery.Document) string {
  span := document.Find(".AHFaub").Find("span")
  return strings.TrimSpace(span.Text())
}

func getAppId(document *goquery.Document) string {
  linkElements := document.Find("link")
  playStoreUrl := getPlayStoreUrl(linkElements)
  parsedUrl, _ := url.Parse(playStoreUrl)
  return parsedUrl.Query().Get("id")
}

func getPlayStoreUrl(linkElements *goquery.Selection) string {
  var playStoreUrl = ""
  linkElements.Each(func(i int, selection *goquery.Selection) {
    attr, exists := selection.Attr("rel")
    if exists && attr == "canonical" {
      url, _ := selection.Attr("href")
      playStoreUrl = url
    }
  })
  return playStoreUrl
}
