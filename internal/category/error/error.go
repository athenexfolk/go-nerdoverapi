package domainerror

import "errors"

var (
	ErrCategoryNotFound      = errors.New("category not found")
	ErrCategoryAlreadyExists = errors.New("category with this ID already exists")
)
