package postgres

import (
	"errors"
	"strings"

	taskapi "github.com/chuksgpfr/task-api"
	"github.com/chuksgpfr/task-api/pkg"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (d *DbService) CreateTask(body *taskapi.CreateTask, user *taskapi.User) (*taskapi.Task, error) {
	var task *taskapi.Task
	err := d.DB.Create(&taskapi.Task{
		UserID:      user.ID,
		Title:       body.Title,
		Status:      strings.ToUpper(body.Status),
		Description: body.Description,
	}).Take(&task).Error

	if err != nil {
		return nil, errors.New("Failed to create task")
	}

	return task, nil
}

func (d *DbService) GetTasks(ctx *gin.Context, user *taskapi.User) ([]*taskapi.Task, int64, error) {
	var count int64
	var tasks []*taskapi.Task

	err := d.DB.Order("created_at desc").Scopes(pkg.Paginate(ctx)).Find(&tasks, &taskapi.Task{UserID: user.ID}).Count(&count).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, nil
		}
		return nil, 0, errors.New("failed to fetch data")
	}

	return tasks, count, nil
}

func (d DbService) GetTask(slug string, user *taskapi.User) (*taskapi.Task, error) {
	var task *taskapi.Task

	err := d.DB.Take(&task, &taskapi.Task{Slug: slug, UserID: user.ID}).Error

	if err != nil {
		return nil, errors.New("This task does not exist")
	}

	return task, nil
}

func (d *DbService) UpdateTask(body *taskapi.CreateTask, user *taskapi.User, slug string) (*taskapi.Task, error) {
	var task *taskapi.Task

	err := d.DB.Take(&task, &taskapi.Task{Slug: slug, UserID: user.ID}).Error

	if err != nil {
		return nil, errors.New("This task does not exist")
	}

	err = d.DB.Model(&taskapi.Task{ID: task.ID}).Updates(&taskapi.Task{
		Description: body.Description,
		Status:      body.Status,
		Title:       body.Title,
	}).Take(&task).Error

	if err != nil {
		return nil, errors.New("This task cannot be updated")
	}

	return task, nil
}

func (d *DbService) CompleteTask(user *taskapi.User, slug string) (*taskapi.Task, error) {
	var task *taskapi.Task

	err := d.DB.Take(&task, &taskapi.Task{Slug: slug, UserID: user.ID}).Error

	if err != nil {
		return nil, errors.New("This task does not exist")
	}

	if task.IsCompleted {
		return nil, errors.New("This task has already been completed")
	}

	err = d.DB.Model(&taskapi.Task{ID: task.ID}).Updates(&taskapi.Task{
		IsCompleted: true,
	}).Take(&task).Error

	if err != nil {
		return nil, errors.New("This task cannot be updated")
	}

	return task, nil
}
