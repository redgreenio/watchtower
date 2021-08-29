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
    Name:            getName(document),
    AppId:           getAppId(document),
    ReleasedOn:      getReleasedOn(document),
    Size:            getSize(document),
    Installs:        getInstalls(document),
    Version:         getVersion(document),
    RequiresAndroid: getRequiresAndroid(document),
    ContentRating:   getContentRating(document),
    OfferedBy:       getOfferedBy(document),
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

func getReleasedOn(document *goquery.Document) string {
  return getValueText("Updated", document)
}

func getPlayStoreUrl(linkElement *goquery.Selection) string {
  playStoreUrl, _ := linkElement.Attr("href")
  return playStoreUrl
}

func getSize(document *goquery.Document) string {
  return getValueText("Size", document)
}

func getInstalls(document *goquery.Document) string {
  return getValueText("Installs", document)
}

func getVersion(document *goquery.Document) string {
  return getValueText("Current Version", document)
}

func getRequiresAndroid(document *goquery.Document) string {
  return getValueText("Requires Android", document)
}

func getContentRating(document *goquery.Document) string {
  selection := getValueSelection("Content Rating", document)
  return selection.Find(".htlgb").Find("div").First().Text()
}

func getOfferedBy(document *goquery.Document) string {
  return getValueText("Offered By", document)
}

func getValueText(title string, document *goquery.Document) string {
  return getValueSelection(title, document).Text()
}

func getValueSelection(title string, document *goquery.Document) *goquery.Selection {
  element := document.Find(".BgcNfc")
  var valueSelection *goquery.Selection = nil
  element.Each(func(i int, selection *goquery.Selection) {
    if selection.Text() == title {
      valueSelection = selection.Siblings().Children().Last()
    }
  })
  return valueSelection
}
