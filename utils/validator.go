package utils

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/hwUltra/fb-tools/xerr"
	"github.com/pkg/errors"
	"net/http"
	"reflect"
)

func ValidatorCheck(r *http.Request, req interface{}) error {
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		return name
	})
	trans, _ := ut.New(zh.New()).GetTranslator("zh")
	validateErr := translations.RegisterDefaultTranslations(validate, trans)
	if validateErr = validate.StructCtx(r.Context(), req); validateErr != nil {
		for _, err := range validateErr.(validator.ValidationErrors) {
			return xerr.NewErrCodeMsg(xerr.REUQES_PARAM_ERROR, errors.New(err.Translate(trans)).Error())
		}
	}
	return nil
}
