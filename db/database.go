package db

import (
	"context"
	"encoding/json"
	"os"
	"strings"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

var Client *firestore.Client

func InitFirestore() {
	projectID := os.Getenv("FIREBASE_PROJECT_ID")
	clientEmail := os.Getenv("FIREBASE_CLIENT_EMAIL")
	privateKey := os.Getenv("FIREBASE_PRIVATE_KEY")

	privateKey = strings.ReplaceAll(privateKey, `\n`, "\n")

	creds := map[string]any{
		"type":                        "service_account",
		"project_id":                  projectID,
		"private_key_id":              "",
		"private_key":                 privateKey,
		"client_email":                clientEmail,
		"client_id":                   "",
		"auth_uri":                    "https://accounts.google.com/o/oauth2/auth",
		"token_uri":                   "https://oauth2.googleapis.com/token",
		"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
		"client_x509_cert_url":        "",
	}

	credsJSON, _ := json.Marshal(creds)

	ctx := context.Background()

	client, err := firestore.NewClient(ctx, projectID, option.WithCredentialsJSON(credsJSON))
	if err != nil {
		panic(err)
	}
	Client = client
}
