package storegcs
// Based on a gcs blog post https://medium.com/wesionary-team/golang-image-upload-with-google-cloud-storage-and-gin-part-2-99f4a642e06a

import (
    "context"
    "io"
    "net/url"
    "giphy-to-gcs/getgiphy"

    "google.golang.org/api/option"
	"cloud.google.com/go/storage"
     log "github.com/sirupsen/logrus"
)

const (
    BUCKET_NAME = "test-logs-dp"
)

var (
	storageClient *storage.Client
)

func ExecuteFileUpload(gifUrl string, id string) {
	var err error
    ctx := context.Background() // What is this? Answer: context package, built-in, Background is default

	storageClient, err = storage.NewClient(ctx, option.WithCredentialsFile("keys.json"))
	if err != nil {
		log.WithFields(log.Fields{ "message": err.Error(), }).Fatal("Cannot create client")
	}

	read, err := getgiphy.GetFile(gifUrl)
	if err != nil {
		log.WithFields(log.Fields{ "message": err.Error(), }).Fatal("Cannot get file")
	}

	defer read.Close()

    log.WithFields(log.Fields{
        "bucket": BUCKET_NAME,
        "id": id,
    }).Info("Uploading to google")

	sw := storageClient.Bucket(BUCKET_NAME).Object(id).NewWriter(ctx)

	if _, err := io.Copy(sw, read); err != nil {
		log.WithFields(log.Fields{ "message": err.Error(), }).Fatal("Cannot copy")
	}

	if err := sw.Close(); err != nil {
		log.WithFields(log.Fields{ "message": err.Error(), }).Fatal("Cannot close writer")
	}

	u, err := url.Parse("/" + BUCKET_NAME + "/" + sw.Attrs().Name)
	if err != nil {
		log.WithFields(log.Fields{ "message": err.Error(), "error":   true, }).Fatal("Cannot parse")
	}

    log.WithFields(log.Fields{
        "gifUrl": gifUrl,
        "pathname": u.EscapedPath(),
    }).Info("Upload Successful")
}
