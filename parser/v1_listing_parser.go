package parser

import (
  "github.com/PuerkitoBio/goquery"
  _ "github.com/PuerkitoBio/goquery"
  "net/url"
  "strings"
  "time"
)

const ReleasedOnDateLayoutV1 string = "January 2, 2006"

type V1ListingParser struct {
}

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

func (parser V1ListingParser) Parse(content string) Release {
  document, _ := goquery.NewDocumentFromReader(strings.NewReader(content))

  return Release{
    Name:            getNameV1(document),
    AppId:           getAppIdV1(document),
    ReleasedOn:      getReleasedOnV1(document),
    Size:            getValueTextV1(sizeTitleText, document),
    Installs:        getValueTextV1(installsTitleText, document),
    Version:         getValueTextV1(versionTitleText, document),
    RequiresAndroid: getValueTextV1(requiresAndroidTitleText, document),
    ContentRating:   getContentRatingV1(document),
    OfferedBy:       getValueTextV1(offeredByTitleText, document),
  }
}

func getNameV1(document *goquery.Document) string {
  span := document.Find(nameClassSelector).Find(htmlSpan)
  return strings.TrimSpace(span.Text())
}

func getAppIdV1(document *goquery.Document) string {
  link := document.Find(appIdSelector)
  playStoreUrl := getPlayStoreUrlV1(link)
  parsedUrl, _ := url.Parse(playStoreUrl)
  return parsedUrl.Query().Get(appIdQueryParameterName)
}

func getReleasedOnV1(document *goquery.Document) time.Time {
  dateText := getValueTextV1(releasedOnTitleText, document)
  date, _ := time.Parse(ReleasedOnDateLayoutV1, dateText)
  return date
}

func getPlayStoreUrlV1(linkElement *goquery.Selection) string {
  playStoreUrl, _ := linkElement.Attr(htmlAttrHref)
  return playStoreUrl
}

func getContentRatingV1(document *goquery.Document) string {
  selection := getValueSelectionV1(contentRatingTitleText, document)
  return selection.Find(contentRatingSelector).Find(htmlDiv).First().Text()
}

func getValueTextV1(title string, document *goquery.Document) string {
  return getValueSelectionV1(title, document).Text()
}

func getValueSelectionV1(title string, document *goquery.Document) *goquery.Selection {
  element := document.Find(valueSelector)
  desiredTitleSelection := element.FilterFunction(func(i int, selection *goquery.Selection) bool {
    return selection.Text() == title
  })
  return desiredTitleSelection.Siblings().Children().Last()
}
