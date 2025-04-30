package lesson

import (
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"nerdoverapi/db"
)

func docRef(id string) *firestore.DocumentRef {
	return db.Client.Collection("lesson").Doc(id)
}

func LessonExists(ctx context.Context, id string) (bool, error) {
	doc, err := docRef(id).Get(ctx)
	if err != nil {
		if isNotFoundErr(err) {
			return false, nil
		}
		return false, err
	}
	return doc.Exists(), nil
}

func isNotFoundErr(err error) bool {
	return status.Code(err) == codes.NotFound
}
