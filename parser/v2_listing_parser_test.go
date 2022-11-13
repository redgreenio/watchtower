package parser

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "time"
)

func TestParseV2Listing(t *testing.T) {
  // given
  content := ReadTestDataFile("v2-listing.html")
  releasedOn, _ := time.Parse(ReleasedOnDateLayoutV2, "29 Sep 2022")
  expected := Release{
    Name:       "Hotstar",
    AppId:      "in.startv.hotstar",
    ReleasedOn: releasedOn,
    Installs:   "500M+",
    OfferedBy:  "Novi Digital",
  }

  // when
  actual := V2ListingParser{}.Parse(content)

  // then
  assert.Equal(t, expected, actual)
}
