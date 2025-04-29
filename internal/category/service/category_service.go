package service

import (
	"context"

	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"nerdoverapi/db"
	"nerdoverapi/internal/category/error"
	"nerdoverapi/internal/category/model"
)

const collectionName = "category"

func CreateCategory(ctx context.Context, newCategory model.Category) (model.Category, error) {
	exists, err := CategoryExists(ctx, newCategory.Slug)
	if err != nil {
		return model.Category{}, err
	}
	if exists {
		return model.Category{}, domainerror.ErrCategoryAlreadyExists
	}

	if _, err := docRef(newCategory.Slug).Set(ctx, newCategory); err != nil {
		return model.Category{}, err
	}
	return newCategory, nil
}

func GetAllCategories(ctx context.Context) ([]model.Category, error) {
	iter := db.Client.Collection(collectionName).Documents(ctx)
	defer iter.Stop()

	var categoryList []model.Category
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var category model.Category
		if err := doc.DataTo(&category); err != nil {
			return nil, err
		}
		categoryList = append(categoryList, category)
	}
	return categoryList, nil
}

func GetCategoryByID(ctx context.Context, id string) (model.Category, error) {
	doc, err := docRef(id).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return model.Category{}, domainerror.ErrCategoryNotFound
		}
		return model.Category{}, err
	}

	var category model.Category
	if err := doc.DataTo(&category); err != nil {
		return model.Category{}, err
	}
	return category, nil
}

func UpdateCategory(ctx context.Context, id string, updatedCategory model.Category) (model.Category, error) {
	exists, err := CategoryExists(ctx, id)
	if err != nil {
		return model.Category{}, err
	}
	if !exists {
		return model.Category{}, domainerror.ErrCategoryNotFound
	}

	updatedCategory.Slug = id
	if _, err := docRef(id).Set(ctx, updatedCategory); err != nil {
		return model.Category{}, err
	}
	return updatedCategory, nil
}

func DeleteCategory(ctx context.Context, id string) (model.Category, error) {
	category, err := GetCategoryByID(ctx, id)
	if err != nil {
		return model.Category{}, err
	}

	if _, err := docRef(id).Delete(ctx); err != nil {
		return model.Category{}, err
	}
	return category, nil
}
