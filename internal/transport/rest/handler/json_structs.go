package handler

type AddCourse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type AddModule struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	OrderID     uint   `json:"order_id"`
	CourseName  string `json:"course_name"`
}

type Delete struct {
	ID uint `json:"id"`
}
