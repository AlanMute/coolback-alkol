package handler

type AddCourse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type AddModule struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	CourseName  string `json:"course_name"`
}

type DeleteCourse struct {
	Name string `json:"name"`
}

type DeleteModule struct {
	Name       string `json:"name"`
	CourseName string `json:"courseName"`
}

type DeleteLesson struct {
	Name       string `json:"name"`
	CourseName string `json:"courseName"`
	ModuleName string `json:"moduleName"`
}
