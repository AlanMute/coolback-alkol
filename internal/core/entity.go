package core

type ModLes struct {
	Module
	Lessons []Lesson
}

type LesMd struct {
	Lesson
	Mdfile []string
}
