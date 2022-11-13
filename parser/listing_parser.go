package parser

type ListingParser interface {
  Parse(content string) Release
}
