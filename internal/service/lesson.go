package service

import (
	"crypto/tls"
	"fmt"
	"os"
	"path/filepath"
	"strings"

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

func (s *LessonService) Add(name string, description string, orderID uint, moduleID uint, content []string) error {
	name = strings.Trim(name, " ")

	description = strings.Trim(description, " ")

	if name == "" {
		name = "New lesson " + fmt.Sprint(orderID)
	}

	lessonID, err := s.repo.Add(name, description, orderID, moduleID)
	if err != nil {
		return err
	}

	filePath := filepath.Join("./lessons", fmt.Sprint(lessonID)+ext)

	if err := pkg.CreateFile(filePath, content); err != nil {
		_, errDel := s.repo.Delete(lessonID)
		if errDel != nil {
			return errDel
		}
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

func (s *LessonService) Get(moduleid int, orderid int) (core.LesMd, error) {
	var lesmd core.LesMd

	lesson, err := s.repo.Get(moduleid, orderid)

	if err != nil {
		return lesmd, err
	}

	path := filepath.Join("lessons", fmt.Sprint(lesson.ID)+ext)

	file, err := pkg.ReadFile(path)

	if err != nil {
		return lesmd, err
	}

	lesmd = core.LesMd{
		Lesson: lesson,
		Mdfile: file,
	}

	return lesmd, nil
}

func (s *LessonService) Put(id int, name string, desc string, orderID uint, content []string) error {
	if len(content) > 0 {
		path := filepath.Join("lessons", fmt.Sprint(id)+ext)
		if err := pkg.UpdateFile(path, content); err != nil {
			return err
		}
	}

	err := s.repo.Put(id, name, desc, orderID)

	if err != nil {
		return err
	}

	return nil
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

	fmt.Println(config.Address, os.Getenv("MAIL_PASSWORD"))
	dialer := gomail.NewDialer("smtp.gmail.com", 587, config.Address, os.Getenv("MAIL_PASSWORD"))

	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := dialer.DialAndSend(message); err != nil {
		return err
	}

	return nil
}
