package parser

import (
  "github.com/PuerkitoBio/goquery"
  "strings"
)

type V2ListingParser struct {
}

func (parser V2ListingParser) Parse(content string) Release {
  document, _ := goquery.NewDocumentFromReader(strings.NewReader(content))

  return Release{
    Name:      getNameV2(document),
    AppId:     getAppIdV2(document),
    Installs:  getInstallsV2(document),
    OfferedBy: getOfferedByV2(document),
  }
}

func getNameV2(document *goquery.Document) string {
  nameSpan := document.Find("div > h1 > span")
  return strings.TrimSpace(nameSpan.Text())
}

func getAppIdV2(document *goquery.Document) string {
  appIdMetaTag := document.Find("meta[itemprop]").First()

  url, _ := appIdMetaTag.Attr("content")
  startIndex := strings.Index(url, "id=") + 3
  endIndex := strings.Index(url, "&")
  return strings.TrimSpace(url[startIndex:endIndex])
}

func getInstallsV2(document *goquery.Document) string {
  firstFlexItem := document.Find("div .wVqUob").First()
  installsNode := firstFlexItem.Next().Children().First()
  return strings.TrimSpace(installsNode.Text())
}

func getOfferedByV2(document *goquery.Document) string {
  offeredBySpan := document.Find("div > a > span").First()
  return strings.TrimSpace(offeredBySpan.Text())
}
