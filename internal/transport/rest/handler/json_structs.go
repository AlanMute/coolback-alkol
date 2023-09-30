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

type Delete struct {
	ID uint `json:"id"`
}

type EdLesson struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	OrderID     int      `json:"orderid"`
	Content     []string `json:"content"`
}
