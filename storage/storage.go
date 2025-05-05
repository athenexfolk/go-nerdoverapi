package storage

import (
	"context"
	"os"

	"cloud.google.com/go/storage"
)

var Bucket *storage.BucketHandle

func InitStorage() {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		panic(err)
	}

	Bucket = client.Bucket(os.Getenv("BUCKET_NAME"))
}
