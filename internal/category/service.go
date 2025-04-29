package category

import (
	"errors"
)

var categories = map[string]Category{
	"a": {ID: "a", Name: "AAA"},
	"b": {ID: "b", Name: "BBB"},
}

func CategoryExists(id string) bool {
	_, exists := categories[id]
	return exists
}

func CreateCategory(newCategory Category) (Category, error) {
	if CategoryExists(newCategory.ID) {
		return Category{}, errors.New("Category with this ID already exists")
	}
	categories[newCategory.ID] = newCategory
	return newCategory, nil
}

func GetAllCategories() []Category {
	categoryList := make([]Category, 0, len(categories))
	for _, category := range categories {
		categoryList = append(categoryList, category)
	}
	return categoryList
}

func GetCategoryByID(id string) (Category, error) {
	category, exists := categories[id]
	if !exists {
		return Category{}, errors.New("Category not found")
	}
	return category, nil
}

func UpdateCategory(id string, updatedCategory Category) (Category, error) {
	if !CategoryExists(id) {
		return Category{}, errors.New("Category not found")
	}
	updatedCategory.ID = id
	categories[id] = updatedCategory
	return updatedCategory, nil
}

func DeleteCategory(id string) (Category, error) {
	if !CategoryExists(id) {
		return Category{}, errors.New("Category not found")
	}
	deletedCategory := categories[id]
	delete(categories, id)
	return deletedCategory, nil
}
