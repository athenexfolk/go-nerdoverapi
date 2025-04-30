package category

import (
	"context"

	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"nerdoverapi/db"
)

const categoryCollectionName = "category"

func CreateCategory(ctx context.Context, newCategory Category) (Category, error) {
	exists, err := CategoryExists(ctx, newCategory.Slug)
	if err != nil {
		return Category{}, err
	}
	if exists {
		return Category{}, ErrCategoryAlreadyExists
	}

	if _, err := docRef(newCategory.Slug).Set(ctx, newCategory); err != nil {
		return Category{}, err
	}
	return newCategory, nil
}

func GetAllCategories(ctx context.Context) ([]Category, error) {
	iter := db.Client.Collection(categoryCollectionName).Documents(ctx)
	defer iter.Stop()

	var categoryList []Category
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var category Category
		if err := doc.DataTo(&category); err != nil {
			return nil, err
		}
		categoryList = append(categoryList, category)
	}
	return categoryList, nil
}

func GetCategoryByID(ctx context.Context, id string) (Category, error) {
	doc, err := docRef(id).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return Category{}, ErrCategoryNotFound
		}
		return Category{}, err
	}

	var category Category
	if err := doc.DataTo(&category); err != nil {
		return Category{}, err
	}
	return category, nil
}

func UpdateCategory(ctx context.Context, id string, updatedCategory Category) (Category, error) {
	exists, err := CategoryExists(ctx, id)
	if err != nil {
		return Category{}, err
	}
	if !exists {
		return Category{}, ErrCategoryNotFound
	}

	updatedCategory.Slug = id
	if _, err := docRef(id).Set(ctx, updatedCategory); err != nil {
		return Category{}, err
	}
	return updatedCategory, nil
}

func DeleteCategory(ctx context.Context, id string) (Category, error) {
	category, err := GetCategoryByID(ctx, id)
	if err != nil {
		return Category{}, err
	}

	if _, err := docRef(id).Delete(ctx); err != nil {
		return Category{}, err
	}
	return category, nil
}
