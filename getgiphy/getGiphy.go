package getgiphy

import (
    "io"
    "io/ioutil" //io in 1.16+
    "net/http"
    "encoding/json"
    "fmt"
    log "github.com/sirupsen/logrus"
)

const (
    GIPHY_API_KEY = "Y4VOEPAxCr4lpdGhtYC5AZmZLV2wP3dF"
    BASE_GIPHY = "https://api.giphy.com/v1"
)

func GetUrl(tag string) (string, string) {
    log.WithFields(log.Fields{"tag":tag}).Info("Tagging with")
    url := fmt.Sprintf("%v/gifs/random?api_key=%v&tag=%v&rating=pg-13", BASE_GIPHY, GIPHY_API_KEY, tag)
    resp, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    //fmt.Println("Body:")
    //fmt.Printf("%s\n", body)
    var treeData interface{}
    json.Unmarshal(body, &treeData)
    rootObj := treeData.(map[string]interface{})
    //fmt.Println(realData["data"]["images"]["preview_webp"])
    dataObj := rootObj["data"].(map[string]interface{})
    imagesObj := dataObj["images"].(map[string]interface{})
    webpObj := imagesObj["preview_webp"].(map[string]interface{})
    webpUrl := webpObj["url"].(string)
    id := dataObj["id"].(string)
    log.Debug(webpUrl)
    return webpUrl, id
}

func GetFile(url string) (read io.ReadCloser, err error) {
    log.WithFields(log.Fields{"url":url}).Info("Fetching image")
    resp, err := http.Get(url)
    if err != nil {
        return resp.Body, err
    }
    return resp.Body, nil
}
