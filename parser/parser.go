package parser

import (
  "github.com/PuerkitoBio/goquery"
  _ "github.com/PuerkitoBio/goquery"
  "net/url"
  "strings"
)

const nameClassSelector = ".AHFaub"
const htmlSpan = "span"

const appIdSelector = "[rel=canonical]"
const appIdQueryParameterName = "id"

const valueSelector = ".BgcNfc"
const releasedOnTitleText = "Updated"
const sizeTitleText = "Size"
const installsTitleText = "Installs"
const versionTitleText = "Current Version"
const requiresAndroidTitleText = "Requires Android"
const offeredByTitleText = "Offered By"

const contentRatingSelector = ".htlgb"
const contentRatingTitleText = "Content Rating"

const htmlAttrHref = "href"
const htmlDiv = "div"

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
  span := document.Find(nameClassSelector).Find(htmlSpan)
  return strings.TrimSpace(span.Text())
}

func getAppId(document *goquery.Document) string {
  link := document.Find(appIdSelector)
  playStoreUrl := getPlayStoreUrl(link)
  parsedUrl, _ := url.Parse(playStoreUrl)
  return parsedUrl.Query().Get(appIdQueryParameterName)
}

func getReleasedOn(document *goquery.Document) string {
  return getValueText(releasedOnTitleText, document)
}

func getPlayStoreUrl(linkElement *goquery.Selection) string {
  playStoreUrl, _ := linkElement.Attr(htmlAttrHref)
  return playStoreUrl
}

func getSize(document *goquery.Document) string {
  return getValueText(sizeTitleText, document)
}

func getInstalls(document *goquery.Document) string {
  return getValueText(installsTitleText, document)
}

func getVersion(document *goquery.Document) string {
  return getValueText(versionTitleText, document)
}

func getRequiresAndroid(document *goquery.Document) string {
  return getValueText(requiresAndroidTitleText, document)
}

func getContentRating(document *goquery.Document) string {
  selection := getValueSelection(contentRatingTitleText, document)
  return selection.Find(contentRatingSelector).Find(htmlDiv).First().Text()
}

func getOfferedBy(document *goquery.Document) string {
  return getValueText(offeredByTitleText, document)
}

func getValueText(title string, document *goquery.Document) string {
  return getValueSelection(title, document).Text()
}

func getValueSelection(title string, document *goquery.Document) *goquery.Selection {
  element := document.Find(valueSelector)
  var valueSelection *goquery.Selection = nil
  element.Each(func(i int, selection *goquery.Selection) {
    if selection.Text() == title {
      valueSelection = selection.Siblings().Children().Last()
    }
  })
  return valueSelection
}
