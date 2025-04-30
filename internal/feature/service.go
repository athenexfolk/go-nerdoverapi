package feature

import (
	"archive/zip"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"time"

	gcs "cloud.google.com/go/storage"
	"google.golang.org/api/iterator"

	// "nerdoverapi/db"
	"nerdoverapi/internal/category"
	"nerdoverapi/internal/lesson"
	"nerdoverapi/storage"
)

func UploadImage(ctx context.Context, file multipart.File, name string, contentType string) (string, error) {
	timestamp := time.Now().UTC().Unix()
	filename := fmt.Sprintf("media/%d_%s", timestamp, name)

	wc := storage.Bucket.Object(filename).NewWriter(ctx)
	wc.ContentType = contentType
	if _, err := io.Copy(wc, file); err != nil {
		return "", err
	}
	if err := wc.Close(); err != nil {
		return "", err
	}

	acl := storage.Bucket.Object(filename).ACL()
	if err := acl.Set(ctx, gcs.AllUsers, gcs.RoleReader); err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://storage.googleapis.com/%s/%s", storage.Bucket.BucketName(), filename)
	return url, nil
}

func GetAllImages(ctx context.Context) ([]string, error) {
	iter := storage.Bucket.Objects(ctx, &gcs.Query{
		Prefix: "media/",
	})

	var urls []string

	for {
		objAttrs, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return urls, err
		}
		url := fmt.Sprintf("https://storage.googleapis.com/%s/%s", storage.Bucket.BucketName(), objAttrs.Name)
		urls = append(urls, url)
	}

	return urls, nil
}

func ExportLesson(ctx context.Context) ([]byte, error) {
	/* Stage 1 : Retrieve data */
	categories, err := category.GetAllCategories(ctx)
	if err != nil {
		return nil, err
	}

	lessons, err := lesson.GetAllLessons(ctx)
	if err != nil {
		return nil, err
	}

	/* Stage 2 : Create menu */
	lessonMap := make(map[string][]MenuLesson)
	for _, lesson := range lessons {
		lessonMap[lesson.CategorySlug] = append(lessonMap[lesson.CategorySlug], MenuLesson{
			Title: lesson.Title,
			Slug:  lesson.Slug,
		})
	}
	var menuList []Menu
	for _, category := range categories {
		menuList = append(menuList, Menu{
			Name:    category.Name,
			Slug:    category.Slug,
			Lessons: lessonMap[category.Slug],
		})
	}

	/* Stage 3 : Convert menu to json */
	jsonMenu, err := json.MarshalIndent(menuList, "", "  ")
	if err != nil {
		return nil, err
	}

	// Create buffer and zip
	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)

	/* stage 4 : Store menu.json file */
	// 4.1 : Create file [menu.json] in zip
	zipFileWriter, err := zipWriter.Create("menu.json")
	if err != nil {
		return nil, err
	}

	// 4.2 : Write menu to menu.json
	_, err = zipFileWriter.Write(jsonMenu)
	if err != nil {
		return nil, err
	}

	/* Stage 5 : Add each lesson to zip */
	for _, lesson := range lessons {
		if lesson.ContentPath == "" {
			continue
		}
		resp, err := http.Get(lesson.ContentPath)
		if err != nil {
			continue
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			continue
		}
		content, err := io.ReadAll(resp.Body)
		if err != nil {
			continue
		}
		contentStr := string(content)
		lesson.Content = &contentStr
		jsonLesson, err := json.MarshalIndent(lesson, "", "  ")
		if err != nil {
			continue
		}

		lessonZipFileWriter, err := zipWriter.Create(lesson.CategorySlug + "." + lesson.Slug + ".json")
		if err != nil {
			continue
		}
		_, err = lessonZipFileWriter.Write(jsonLesson)
		if err != nil {
			continue
		}
	}

	err = zipWriter.Close()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
