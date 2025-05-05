package db

import (
	"context"
	"os"

	"cloud.google.com/go/firestore"
)

var Client *firestore.Client

func InitFirestore() {
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		panic(err)
	}
	Client = client
}
