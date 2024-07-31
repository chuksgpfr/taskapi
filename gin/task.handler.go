package gin

import (
	"net/http"

	taskapi "github.com/chuksgpfr/task-api"
	"github.com/chuksgpfr/task-api/pkg"
	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	TaskService taskapi.TaskService
	Config      pkg.Configuration
}

func (w *TaskHandler) getUserFromContext(ctx *gin.Context) *taskapi.User {
	userStr, ok := ctx.Get("user")

	if !ok {
		return nil
	}

	user := userStr.(*taskapi.User)

	return user
}

func (w *TaskHandler) CreateTask(ctx *gin.Context) {
	var body *taskapi.CreateTask

	ctx.ShouldBind(&body)

	err := ValidateBody(body, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err.Error()))
		return
	}

	user := w.getUserFromContext(ctx)
	if user == nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse("An invalid user was passed"))
		return
	}

	task, err := w.TaskService.CreateTask(body, user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err.Error()))
		return
	}

	resp := map[string]interface{}{
		"task": task,
	}

	ctx.JSON(http.StatusOK, SuccessResponse("Task has been set", resp))
}

func (t *TaskHandler) GetTasks(ctx *gin.Context) {
	limit, page := pkg.GetPageQueries(ctx)

	user := t.getUserFromContext(ctx)
	if user == nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse("An invalid user was passed"))
		return
	}

	tasks, count, err := t.TaskService.GetTasks(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusOK, SuccessResponse("Ok", nil))
		return
	}

	nextPage, currentPage, total := pkg.PaginationDetails(uint64(limit), uint64(page), uint64(count))

	resp := map[string]interface{}{
		"tasks":       tasks,
		"total":       total,
		"nextPage":    nextPage,
		"currentPage": currentPage,
	}

	ctx.JSON(http.StatusOK, SuccessResponse("Ok", resp))
	return
}

func (t *TaskHandler) GetTask(ctx *gin.Context) {
	user := t.getUserFromContext(ctx)
	if user == nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse("An invalid user was passed"))
		return
	}
	slug, _ := ctx.Params.Get("slug")

	task, err := t.TaskService.GetTask(slug, user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err.Error()))
		return
	}

	resp := map[string]interface{}{
		"task": task,
	}

	ctx.JSON(http.StatusOK, SuccessResponse("Ok", resp))
	return
}

func (w *TaskHandler) UpdateTask(ctx *gin.Context) {
	var body *taskapi.CreateTask

	ctx.ShouldBind(&body)

	err := ValidateBody(body, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err.Error()))
		return
	}

	user := w.getUserFromContext(ctx)
	if user == nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse("An invalid user was passed"))
		return
	}
	slug, _ := ctx.Params.Get("slug")

	task, err := w.TaskService.UpdateTask(body, user, slug)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err.Error()))
		return
	}

	resp := map[string]interface{}{
		"task": task,
	}

	ctx.JSON(http.StatusOK, SuccessResponse("Task has been updated", resp))
}

func (w *TaskHandler) CompleteTask(ctx *gin.Context) {
	user := w.getUserFromContext(ctx)
	if user == nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse("An invalid user was passed"))
		return
	}
	slug, _ := ctx.Params.Get("slug")

	task, err := w.TaskService.CompleteTask(user, slug)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err.Error()))
		return
	}

	resp := map[string]interface{}{
		"task": task,
	}

	ctx.JSON(http.StatusOK, SuccessResponse("Task has been completed", resp))
}
