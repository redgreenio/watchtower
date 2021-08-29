package parser

import (
  "github.com/stretchr/testify/assert"
  "io/ioutil"
  "testing"
  "time"
)

func TestParse(t *testing.T) {
  // given
  content := ReadTestDataFile("dunzo-delivery-partner.html")
  releasedOn, _ := time.Parse(DateLayout, "July 19, 2021")
  expected := &PlayStoreAppListing{
    Name:            "Dunzo Delivery Partner",
    AppId:           "runner.dunzo.com.dunzo_runner",
    ReleasedOn:      releasedOn,
    Size:            "29M",
    Installs:        "100,000+",
    Version:         "3.22.0.0",
    RequiresAndroid: "5.0 and up",
    ContentRating:   "Everyone",
    OfferedBy:       "Dunzo Digital",
  }

  // when
  actual := Parse(content)

  // then
  assert.Equal(t, expected, actual)
}

func ReadTestDataFile(filename string) string {
  bytes, _ := ioutil.ReadFile("../testdata/valid-play-store-listing-v1/" + filename)
  return string(bytes)
}
