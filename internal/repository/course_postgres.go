package repository

import (
	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/jinzhu/gorm"
)

type CoursePostgres struct {
	db *gorm.DB
}

func NewCoursePostgres(db *gorm.DB) *CoursePostgres {
	return &CoursePostgres{db: db}
}

// Входные данные тут не нужны так как все хранится в контексте будет

// func AddCourse(repo repository.Repository) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var info AddCourseModule

// 		if err := c.ShouldBindJSON(&info); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		path := uniqueFolder("course", "/courses")
// 		slice := strings.Split(path, "/")
// 		folderName := "/" + slice[len(slice)-1]

// 		if err := os.Mkdir("./"+path, os.ModePerm); err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}

// 		repo.AddCourse(info.Name, info.Description, folderName)
// 		c.Status(http.StatusOK)
// 	}
// }

func (r *CoursePostgres) Add(name string, description string, folderName string) error {
	newCourse := core.Course{
		Name:        name,
		Description: description,
		NameFolder:  folderName,
	}

	if result := r.db.Create(&newCourse); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *CoursePostgres) GetByName(name string) ([]core.Course, error) {

	var courses []core.Course
	r.db.Where("name ILIKE ?", "%"+name+"%").Find(&courses)

	return courses, nil
}

func (r *CoursePostgres) GetAll() ([]core.Course, error) {

	var courses []core.Course
	r.db.Find(&courses)
	return courses, nil
}
