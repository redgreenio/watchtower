package parser

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestParseV2Listing(t *testing.T) {
  // given
  content := ReadTestDataFile("v2-listing.html")
  expected := Release{
    Name:     "Hotstar",
    AppId:    "in.startv.hotstar",
    Installs: "500M+",
  }

  // when
  actual := V2ListingParser{}.Parse(content)

  // then
  assert.Equal(t, expected, actual)
}
