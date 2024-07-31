package gin

import (
	"net/http"
	"time"

	taskapi "github.com/chuksgpfr/task-api"
	"github.com/chuksgpfr/task-api/pkg"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService taskapi.UserService
	Config      pkg.Configuration
}

func (u *UserHandler) Register(ctx *gin.Context) {
	var body *taskapi.RegisterParam

	ctx.ShouldBind(&body)

	err := ValidateBody(body, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err.Error()))
		return
	}

	user, err := u.UserService.Register(body)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err.Error()))
		return
	}

	expiry := time.Now().UTC().Add(24 * time.Hour)

	claims := pkg.TokenClaims{
		ID:        user.ID,
		ExpiresAt: expiry.Unix(),
	}

	jwtToken, err := pkg.GenerateToken(u.Config.LoginSymmetricKey, claims)

	resp := map[string]interface{}{
		"user":  user,
		"token": jwtToken,
	}

	ctx.JSON(http.StatusOK, SuccessResponse("Welcome to task buddy", resp))
	return
}

func (u *UserHandler) Login(ctx *gin.Context) {
	var body *taskapi.LoginParam

	ctx.ShouldBind(&body)

	err := ValidateBody(body, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err.Error()))
		return
	}

	user, err := u.UserService.Login(body)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err.Error()))
		return
	}

	expiry := time.Now().UTC().Add(24 * time.Hour)

	claims := pkg.TokenClaims{
		ID:        user.ID,
		ExpiresAt: expiry.Unix(),
	}

	jwtToken, err := pkg.GenerateToken(u.Config.LoginSymmetricKey, claims)

	resp := map[string]interface{}{
		"user":  user,
		"token": jwtToken,
	}

	ctx.JSON(http.StatusOK, SuccessResponse("Welcome to task buddy", resp))
	return
}
