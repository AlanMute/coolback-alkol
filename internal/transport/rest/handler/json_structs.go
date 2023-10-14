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

type EdLesson struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	OrderID     uint     `json:"orderid"`
	Content     []string `json:"content"`
}

type EdModule struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	OrderID     uint   `json:"orderid"`
}

type SignInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Refresh struct {
	RefreshToken string `json:"refresh_token"`
}
