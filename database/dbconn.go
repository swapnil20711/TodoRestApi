package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Name   string `json:"name"`
	Status string `json:"status"`
}

const (
	host   = "localhost"
	port   = 5432
	user   = "swapnil"
	dbname = "todo-list-api"
)

var DB *gorm.DB

func CreateTask(name string, status string) (Task, error) {
	var newTask = Task{Name: name, Status: status}

	result := DB.Create(&newTask)

	if result.Error != nil {
		return newTask, result.Error
	}

	return newTask, nil
}

func GetAllTasks() ([]Task, error) {
	var tasks []Task
	result := DB.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func GetById(id string) (Task, error) {
	var task Task
	result := DB.Where("ID = ? ", id).First(&task)

	if result.Error != nil {
		return task, result.Error
	}
	return task, nil
}

func InitDB() error {
	var err error
	errGoDotEnv := godotenv.Load()
	if errGoDotEnv != nil {
		log.Fatal("Error loading .env file")
	}
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		return fmt.Errorf("DB_PASSWORD is not set")
	}
	var dsn = fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai", host, port, user, password, dbname)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB.AutoMigrate(&Task{})
	return nil
}
