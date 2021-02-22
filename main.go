package main

import (
    "os"
    "giphy-to-gcs/storegcs"
    "giphy-to-gcs/getgiphy"
    log "github.com/sirupsen/logrus"
)

func main() {
    log.SetLevel(log.DebugLevel)
    tag := "cow"
    if len(os.Args) > 1 {
        tag = os.Args[1]
    }
    url, id := getgiphy.GetUrl(tag)
    storegcs.ExecuteFileUpload(url, id)
}
