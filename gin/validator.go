package gin

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translator "github.com/go-playground/validator/v10/translations/en"
)

var (
	validate = validator.New()
)

func translator() ut.Translator {
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	en_translator.RegisterDefaultTranslations(validate, trans)
	return trans
}

func ValidateBody(body interface{}, ctx *gin.Context) error {
	// trans := translator()
	err := validate.Struct(body)
	if err != nil {
		validatorErrors := err.(validator.ValidationErrors)
		for _, e := range validatorErrors {
			err := matchErrorMsg(e)
			// ctx.JSON(http.StatusBadRequest, ErrorResponse(matchErrorMsg(e)))
			return fmt.Errorf(err)
		}
	}

	return nil
}

func matchErrorMsg(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", strings.ToLower(e.Field()))
	}
	return fmt.Sprintf("%s is invalid", strings.ToLower(e.Field()))
}

func MatchError(err error, ctx *gin.Context) {
	for _, fieldErr := range err.(validator.ValidationErrors) {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(matchErrorMsg(fieldErr)))
		return
	}
}
