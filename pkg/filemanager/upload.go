package filemanager

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/api/option"
	"io"
	"net/url"
	"peanut/config"
	"peanut/pkg/ary"
	"strings"
	"time"
)

func CheckExtensionAvailable(ext string, listExt []string) bool {
	ext = strings.ToLower(ext)
	return ary.InArray(ext, listExt)
}

func UploadGgStorage(fileContent io.Reader, nameFile string) (string, error) {
	// defines Context
	ctx := context.Background()

	// new name to upload
	if nameFile == "" {
		nameFile = "_NoName"
	}
	nameUpload := uuid.NewString() + nameFile

	// bucket name to upload
	bucket := config.BucketUpload
	// Creating a Client with file json Credentials
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(config.GgStorageCredential))
	if err != nil {
		return "", fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	// Upload an object with storage.Writer.
	wc := client.Bucket(bucket).Object(nameUpload).NewWriter(ctx)
	if _, err = io.Copy(wc, fileContent); err != nil {
		return "", fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("Writer.Close: %v", err)
	}

	// get public url image
	url, err := url.Parse(config.PublicUrlGgStorage + "/" + bucket + "/" + wc.Attrs().Name)
	if err != nil {
		return "", err
	}
	return url.String(), nil
}
