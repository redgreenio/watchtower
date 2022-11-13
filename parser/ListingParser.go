package parser

import (
  "github.com/PuerkitoBio/goquery"
  "time"
)

type ListingParser interface {
  getName(document *goquery.Document) string
  getAppId(document *goquery.Document) string
  getReleasedOn(document *goquery.Document) time.Time
  getPlayStoreUrl(linkElement *goquery.Selection) string
  getContentRating(document *goquery.Document) string
  getValueText(title string, document *goquery.Document) string
}
