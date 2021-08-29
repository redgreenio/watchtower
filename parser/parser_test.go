package parser

import (
  "github.com/stretchr/testify/assert"
  "io/ioutil"
  "testing"
)

func TestParse(t *testing.T) {
  // given
  content := ReadTestDataFile("dunzo-delivery-partner.html")

  // when
  actual := Parse(content)

  // then
  expected := &PlayStoreAppListing{
    Name: "Dunzo Delivery Partner",
    AppId: "runner.dunzo.com.dunzo_runner",
  }
  assert.Equal(t, expected, actual)
}

func ReadTestDataFile(filename string) string {
  bytes, _ := ioutil.ReadFile("../testdata/valid-play-store-listing-v1/" + filename)
  return string(bytes)
}
