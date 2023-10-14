package handler

type AddCourse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type AddModule struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	OrderID     uint   `json:"order_id"`
	CourseID    uint   `json:"course_id"`
}

type AddLesson struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	OrderID     uint     `json:"order_id"`
	ModuleID    uint     `json:"module_id"`
	Content     []string `json:"content"`
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
