package service

import (
	"crypto/tls"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"

	"github.com/KrizzMU/coolback-alkol/internal/config/emailConf"
	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/KrizzMU/coolback-alkol/internal/repository"
	"github.com/KrizzMU/coolback-alkol/pkg"
	"gopkg.in/gomail.v2"
)

const (
	ext string = ".md"
)

type LessonService struct {
	repo repository.Lesson
}

func NewLessonService(repo repository.Lesson) *LessonService {
	return &LessonService{repo: repo}
}

func (s *LessonService) Add(file multipart.File, fileName string, name string, description string, orderID uint, moduleName string, courseName string) error {
	fileID, err := s.repo.Add(name, description, orderID, courseName, moduleName)
	if err != nil {
		return err
	}

	filePath := filepath.Join("./lessons", strconv.FormatUint(uint64(fileID), 10)+ext)

	if err := pkg.CreateFile(file, filePath); err != nil {
		return err
	}

	return nil
}

func (s *LessonService) Delete(id uint) error {
	fileID, err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	fileName := fileID + ext

	filePath := filepath.Join("./lessons", fileName)

	if err := os.Remove(filePath); !os.IsNotExist(err) {
		fmt.Printf("err = %e", err)
		return err
	}

	return nil
}

func (s *LessonService) Get(course string, module string, lesson string) (core.LesMd, error) {
	var lesmd core.LesMd

	path := filepath.Join("./courses", course, module, lesson+".md")

	strFile, err := pkg.ReadFile(path)

	if err != nil {
		return lesmd, err
	}

	path = "courses\\" + course + "\\" + module + "\\" + lesson + ".md"

	lesmd, err = s.repo.Get(path, strFile)

	if err != nil {
		return lesmd, err
	}

	return lesmd, nil
}

func (s *LessonService) SendTrialLesson(email string) error {

	// check, err := regexp.Match(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, []byte(email))
	// if err != nil {
	// 	return err
	// } else if !check {
	// 	return fmt.Errorf("wrong email format")
	// }

	// fmt.Println(check, err)

	absolutePath, err := os.Getwd()
	if err != nil {
		return err
	}

	filePath := "/static/letters/trial.pdf"

	pdfFilePath := filepath.Join(absolutePath, filePath)

	if err := s.repo.SendTrialLesson(email); err != nil {
		return err
	}

	if err := sendViaMailRu(email, pdfFilePath); err != nil {
		return err
	}

	return nil
}

func sendViaMailRu(email string, path string) error {
	config := emailConf.GetEmailConfig()

	message := gomail.NewMessage()
	message.SetHeader("From", config.Address)
	message.SetHeader("To", email)
	message.SetHeader("Subject", "Тестовое сообщение")
	message.SetBody("text/plain", "Тестовое сообщение через Golang с файлом")
	message.Attach(path)

	dialer := gomail.NewDialer("smtp.mail.ru", 465, config.Address, config.Password)

	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := dialer.DialAndSend(message); err != nil {
		return err
	}

	return nil
}
