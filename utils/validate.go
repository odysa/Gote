package utils

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"strings"
)

func DefaultGetValidParams(c *gin.Context, params interface{}) error {
	if err := c.ShouldBind(params); err != nil {
		return err
	}

	valid, err := GetValidator(c)
	if err != nil {
		return err
	}

	trans, err := GetTranslation(c)
	if err != nil {
		return err
	}

	err = valid.Struct(params)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		sliceErrs := []string{}
		for _, e := range errs {
			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		return errors.New(strings.Join(sliceErrs, ","))
	}

	return nil
}

func GetValidator(c *gin.Context) (*validator.Validate, error) {
	val, ok := c.Get(ValidatorKey)
	if !ok {
		return nil, errors.New("validator not set")
	}
	v, ok := val.(*validator.Validate)

	if !ok {
		return nil, errors.New("failed to get validator")
	}

	return v, nil
}

func GetTranslation(c *gin.Context) (ut.Translator, error) {
	trans, ok := c.Get(TranslatorKey)
	if !ok {
		return nil, errors.New("translator not set")
	}
	translator, ok := trans.(ut.Translator)
	if !ok {
		return nil, errors.New("failed to get a translator")
	}
	return translator, nil
}
