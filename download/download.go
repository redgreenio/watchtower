package download

import (
"fmt"
  "io/ioutil"
  "log"
"net/http"
)
import _ "net/http"

func Download(appId string) string {
  resp, err := http.Get("https://play.google.com/store/apps/details?id=" + appId)
  if err != nil {
    log.Fatalln(fmt.Sprintf("Uh oh! unable to download listing for '%s'.", appId))
  }
  content, _ := ioutil.ReadAll(resp.Body)
  _ = resp.Body.Close()

  return string(content)
}
