package feature

type Menu struct {
	Name    string
	Slug    string
	Lessons []MenuLesson
}

type MenuLesson struct {
	Title string
	Slug  string
}
