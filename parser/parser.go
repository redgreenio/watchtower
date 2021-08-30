package parser

import (
  "github.com/PuerkitoBio/goquery"
  _ "github.com/PuerkitoBio/goquery"
  "net/url"
  "strings"
  "time"
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

func Parse(content string) *Release {
  document, _ := goquery.NewDocumentFromReader(strings.NewReader(content))

  return &Release{
    Name:            getName(document),
    AppId:           getAppId(document),
    ReleasedOn:      getReleasedOn(document),
    Size:            getValueText(sizeTitleText, document),
    Installs:        getValueText(installsTitleText, document),
    Version:         getValueText(versionTitleText, document),
    RequiresAndroid: getValueText(requiresAndroidTitleText, document),
    ContentRating:   getContentRating(document),
    OfferedBy:       getValueText(offeredByTitleText, document),
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

func getReleasedOn(document *goquery.Document) time.Time {
  dateText := getValueText(releasedOnTitleText, document)
  date, _ := time.Parse(ReleasedOnDateLayout, dateText)
  return date
}

func getPlayStoreUrl(linkElement *goquery.Selection) string {
  playStoreUrl, _ := linkElement.Attr(htmlAttrHref)
  return playStoreUrl
}

func getContentRating(document *goquery.Document) string {
  selection := getValueSelection(contentRatingTitleText, document)
  return selection.Find(contentRatingSelector).Find(htmlDiv).First().Text()
}

func getValueText(title string, document *goquery.Document) string {
  return getValueSelection(title, document).Text()
}

func getValueSelection(title string, document *goquery.Document) *goquery.Selection {
  element := document.Find(valueSelector)
  desiredTitleSelection := element.FilterFunction(func(i int, selection *goquery.Selection) bool {
    return selection.Text() == title
  })
  return desiredTitleSelection.Siblings().Children().Last()
}
