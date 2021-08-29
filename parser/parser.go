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
    Name:     getName(document),
    AppId:    getAppId(document),
    Size:     getSize(document),
    Installs: getInstalls(document),
    Version: getVersion(document),
    RequiresAndroid: getRequiresAndroid(document),
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

func getSize(document *goquery.Document) string {
  return getValue("Size", document)
}

func getInstalls(document *goquery.Document) string {
  return getValue("Installs", document)
}

func getVersion(document *goquery.Document) string {
  return getValue("Current Version", document)
}

func getRequiresAndroid(document *goquery.Document) string {
  return getValue("Requires Android", document)
}

func getValue(title string, document *goquery.Document) string {
  element := document.Find(".BgcNfc")
  size := ""
  element.Each(func(i int, selection *goquery.Selection) {
    if selection.Text() == title {
      size = selection.Siblings().Children().Last().Text()
    }
  })
  return size
}
