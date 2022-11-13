package cmd

import (
  "fmt"
  "github.com/fatih/color"
  "github.com/spf13/cobra"
  "sync"
  "time"
  "watchtower/app"
  "watchtower/database"
  "watchtower/download"
  "watchtower/parser"
  "watchtower/release"
)

var (
  counter   int64 = 1
  mutex     sync.Mutex
  waitGroup sync.WaitGroup
  dbMutex   sync.Mutex
)

var red = color.New(color.FgRed)

var observeCmd = &cobra.Command{
  Use: "observe",
  Run: func(cmd *cobra.Command, args []string) {
    start := time.Now()
    database := database.InitDb(database.DbPath)
    appsRepository := app.DefaultAppsRepository{Database: database}
    releasesRepository := release.DefaultReleasesRepository{Database: database}

    apps := appsRepository.List()
    appsCount := len(apps)

    waitGroup.Add(len(apps))

    for _, app := range apps {
      go downloadAndSaveReleaseInformation(app, appsCount, releasesRepository, &waitGroup)
    }

    waitGroup.Wait()
    elapsed := time.Since(start)
    println(fmt.Sprintf("~ Observation took %dms ~", elapsed.Milliseconds()))
  },
}

func downloadAndSaveReleaseInformation(
  app app.App,
  appsCount int,
  releasesRepository release.ReleasesRepository,
  waitGroup *sync.WaitGroup,
) {
  var count int64
  mutex.Lock()
  count = counter
  println(fmt.Sprintf("[%d/%d] Downloading release info for '%s'", count, appsCount, app.AppId))
  counter++
  mutex.Unlock()

  content, err := download.Download(app.AppId, app.Country)
  if err != nil {
    println(fmt.Sprintf("[%d/%d] Download failed for '%s'", count, appsCount, app.AppId))
  } else {
    appRelease := parser.V1ListingParser{}.Parse(content)
    if appRelease.AppId == "" {
      _, _ = red.Printf("Unable to parse HTML for '%s'", app.AppId)
      println()
    } else {
      dbMutex.Lock()
      releasesRepository.Insert(appRelease)
      dbMutex.Unlock()
    }
  }
  defer waitGroup.Done()
}
