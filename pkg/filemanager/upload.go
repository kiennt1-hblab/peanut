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

func defineGgStorage(ctx context.Context) (*storage.Client, error) {
	// Creating a Client with file json Credentials
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(config.GgStorageCredential))
	if err != nil {
		return nil, fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	_, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	return client, nil
}

func UploadGgStorage(fileContent io.Reader, nameFile string) (string, error) {
	// defines Context
	ctx := context.Background()

	client, err := defineGgStorage(ctx)
	if err != nil {
		return "", err
	}

	// bucket name to upload
	bucket := config.BucketUpload

	// new name to upload
	if nameFile == "" {
		nameFile = "_NoName"
	}
	nameUpload := uuid.NewString() + nameFile

	// Upload an object with storage.Writer.
	wc := client.Bucket(bucket).Object(nameUpload).NewWriter(ctx)
	if _, err = io.Copy(wc, fileContent); err != nil {
		return "", fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("Writer.Close: %v", err)
	}

	// get public url image
	urlReturn, err := url.Parse(config.PublicUrlGgStorage + "/" + bucket + "/" + wc.Attrs().Name)
	if err != nil {
		return "", err
	}
	return urlReturn.String(), nil
}

func DeleteGCS(nameFile string) error {
	// defines Context
	ctx := context.Background()

	client, err := defineGgStorage(ctx)
	if err != nil {
		return err
	}

	// bucket name to upload
	bucket := config.BucketUpload

	// delete file
	err = client.Bucket(bucket).Object(nameFile).Delete(ctx)
	if err != nil {
		return fmt.Errorf("Object(%q).Delete: %v", nameFile, err)
	}

	return nil
}
