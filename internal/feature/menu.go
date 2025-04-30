package feature

type Menu struct {
	Name    string       `json:"name"`
	Slug    string       `json:"slug"`
	Lessons []MenuLesson `json:"lessons"`
}

type MenuLesson struct {
	Title string `json:"title"`
	Slug  string `json:"slug"`
}
