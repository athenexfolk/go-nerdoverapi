package lesson

import "errors"

var (
	ErrLessonNotFound      = errors.New("lesson not found")
	ErrLessonAlreadyExists = errors.New("lesson with this ID already exists")
)
