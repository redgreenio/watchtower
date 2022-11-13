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
    Name:     parser.getName(document),
    AppId:    parser.getAppId(document),
    Installs: parser.getInstalls(document),
  }
}

func (parser V2ListingParser) getName(document *goquery.Document) string {
  nameSpan := document.Find("div > h1 > span")
  return strings.TrimSpace(nameSpan.Text())
}

func (parser V2ListingParser) getAppId(document *goquery.Document) string {
  appIdMetaTag := document.Find("meta[itemprop]").First()

  url, _ := appIdMetaTag.Attr("content")
  startIndex := strings.Index(url, "id=") + 3
  endIndex := strings.Index(url, "&")
  return strings.TrimSpace(url[startIndex:endIndex])
}

func (parser V2ListingParser) getInstalls(document *goquery.Document) string {
  firstFlexItem := document.Find("div .wVqUob").First()
  installsNode := firstFlexItem.Next().Children().First()
  return strings.TrimSpace(installsNode.Text())
}
