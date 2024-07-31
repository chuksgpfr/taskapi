package taskapi

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreateTask struct {
	Title       string `json:"title" validate:"required"`
	Status      string `json:"status" validate:"required,oneof=pending published"`
	Description string `json:"description" validate:"required"`
}

type Task struct {
	gorm.Model
	ID          uint   `json:"id" gorm:"primaryKey"`
	UserID      uint   `json:"userID"`
	User        User   `json:"user" gorm:"foreignKey:UserID"`
	Title       string `json:"title"`
	Status      string `json:"status"`
	Slug        string `json:"slug"`
	IsCompleted bool   `json:"isCompleted" gorm:"default:false"`
	Description string `json:"description"`
}

type TaskService interface {
	CreateTask(body *CreateTask, user *User) (*Task, error)
	GetTasks(ctx *gin.Context, user *User) ([]*Task, int64, error)
	GetTask(slug string, user *User) (*Task, error)
	UpdateTask(body *CreateTask, user *User, slug string) (*Task, error)
	CompleteTask(user *User, slug string) (*Task, error)
}

func (t *Task) BeforeCreate(tx *gorm.DB) (err error) {
	t.Slug = uuid.New().String()
	t.IsCompleted = false
	return
}
