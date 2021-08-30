package download

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
)
import _ "net/http"

func Download(appId, country string) (body string, downloadError error) {
  var playStoreUrl string
  if country != "" {
    playStoreUrl = "https://play.google.com/store/apps/details?id=" + appId + "&gl=" + country
  } else {
    playStoreUrl = "https://play.google.com/store/apps/details?id=" + appId
  }

  resp, err := http.Get(playStoreUrl)
  if err != nil {
    log.Fatalln(fmt.Sprintf("Uh oh! unable to download listing for '%s'.", appId))
    return "", err
  }
  content, _ := ioutil.ReadAll(resp.Body)
  _ = resp.Body.Close()

  return string(content), nil
}
