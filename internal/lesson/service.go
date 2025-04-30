package lesson

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	gcs "cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"nerdoverapi/db"
	"nerdoverapi/internal/category"
	"nerdoverapi/storage"
)

const lessonCollectionName = "lesson"

func CreateLesson(ctx context.Context, dto CreateLessonDto) (Lesson, error) {
	newLesson := Lesson{
		Title:        dto.Title,
		Slug:         dto.Slug,
		CategoryName: dto.CategoryName,
		CategorySlug: dto.CategorySlug,
		Cover:        dto.Cover,
	}

	/* Stage 1 : Check if requested lesson has valid category */
	exist, err := category.CategoryExists(ctx, newLesson.CategorySlug)
	if err != nil {
		return Lesson{}, err
	}
	if !exist {
		return Lesson{}, category.ErrCategoryNotFound
	}

	/* Stage 2 : Check if requested lesson is not duplicate with exist data */
	exists, err := LessonExists(ctx, newLesson.Slug)
	if err != nil {
		return Lesson{}, err
	}
	if exists {
		return Lesson{}, ErrLessonAlreadyExists
	}

	/* Stage 3 : Create MD file and upload to GCS */
	filename := fmt.Sprintf("content/%s.%s.md", newLesson.CategorySlug, newLesson.Slug)

	wc := storage.Bucket.Object(filename).NewWriter(ctx)
	wc.ContentType = "text/markdown; charset=utf-8"
	wc.CacheControl = "public, max-age=3600"
	content := fmt.Appendf(nil, "# %s", newLesson.Title)
	_, err = io.Copy(wc, bytes.NewReader(content))
	if err != nil {
		wc.Close()
		return Lesson{}, err
	}
	if err := wc.Close(); err != nil {
		return Lesson{}, err
	}

	/* Stage 4 : Make file public readable */
	acl := storage.Bucket.Object(filename).ACL()
	if err := acl.Set(ctx, gcs.AllUsers, gcs.RoleReader); err != nil {
		return Lesson{}, err
	}

	/* Stage 5 : Set content URL to MD file */
	newLesson.ContentPath = fmt.Sprintf("https://storage.googleapis.com/%s/%s", storage.Bucket.BucketName(), filename)

	/* Stage 6 : Add record to DB */
	if _, err := docRef(newLesson.Slug).Set(ctx, newLesson); err != nil {
		return Lesson{}, err
	}

	return newLesson, nil
}

func GetAllLessons(ctx context.Context) ([]Lesson, error) {
	iter := db.Client.Collection(lessonCollectionName).Documents(ctx)
	defer iter.Stop()

	var lessonList []Lesson
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var lesson Lesson
		if err := doc.DataTo(&lesson); err != nil {
			return nil, err
		}
		lessonList = append(lessonList, lesson)
	}
	return lessonList, nil
}

func GetLessonByID(ctx context.Context, id string) (Lesson, error) {
	doc, err := docRef(id).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return Lesson{}, ErrLessonNotFound
		}
		return Lesson{}, err
	}

	var lesson Lesson
	if err := doc.DataTo(&lesson); err != nil {
		return Lesson{}, err
	}

	if lesson.ContentPath != "" {
		resp, err := http.Get(lesson.ContentPath)
		if err != nil {
			return Lesson{}, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return Lesson{}, err
		}

		content, err := io.ReadAll(resp.Body)
		if err != nil {
			return Lesson{}, err
		}

		contentStr := string(content)
		lesson.Content = &contentStr
	}

	return lesson, nil
}

func UpdateLesson(ctx context.Context, id string, dto UpdateLessonDto) (Lesson, error) {
	exists, err := LessonExists(ctx, id)
	if err != nil {
		return Lesson{}, err
	}
	if !exists {
		return Lesson{}, ErrLessonNotFound
	}

	docSnap, err := docRef(id).Get(ctx)
	if err != nil {
		return Lesson{}, err
	}

	var existingLesson Lesson
	if err := docSnap.DataTo(&existingLesson); err != nil {
		return Lesson{}, err
	}

	if dto.Title != nil {
		existingLesson.Title = *dto.Title
	}
	if dto.Cover != nil {
		existingLesson.Cover = dto.Cover
	}

	if _, err := docRef(id).Set(ctx, existingLesson); err != nil {
		return Lesson{}, err
	}
	return existingLesson, nil
}

func UpdateContent(ctx context.Context, id string, dto UpdateContentDto) (Lesson, error) {
	exists, err := LessonExists(ctx, id)
	if err != nil {
		return Lesson{}, err
	}
	if !exists {
		return Lesson{}, ErrLessonNotFound
	}

	docSnap, err := docRef(id).Get(ctx)
	if err != nil {
		return Lesson{}, err
	}

	var existingLesson Lesson
	if err := docSnap.DataTo(&existingLesson); err != nil {
		return Lesson{}, err
	}

	filename := fmt.Sprintf("content/%s.%s.md", existingLesson.CategorySlug, existingLesson.Slug)

	wc := storage.Bucket.Object(filename).NewWriter(ctx)
	wc.ContentType = "text/markdown; charset=utf-8"
	wc.CacheControl = "public, max-age=3600"
	content := []byte(dto.Content)
	_, err = io.Copy(wc, bytes.NewReader(content))
	if err != nil {
		wc.Close()
		return Lesson{}, err
	}
	if err := wc.Close(); err != nil {
		return Lesson{}, err
	}

	acl := storage.Bucket.Object(filename).ACL()
	if err := acl.Set(ctx, gcs.AllUsers, gcs.RoleReader); err != nil {
		return Lesson{}, err
	}

	return existingLesson, nil
}

func DeleteLesson(ctx context.Context, id string) (Lesson, error) {
	lesson, err := GetLessonByID(ctx, id)
	if err != nil {
		return Lesson{}, err
	}

	if _, err := docRef(id).Delete(ctx); err != nil {
		return Lesson{}, err
	}
	return lesson, nil
}
