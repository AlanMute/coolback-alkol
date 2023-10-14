package core

type ModLes struct {
	Module
	Lessons []Lesson
}

type LesMd struct {
	Lesson
	Mdfile []string
}

type CourseContent struct {
	Course
	Modules []ModLes
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}
