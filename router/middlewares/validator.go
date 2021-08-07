package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"github.com/odysa/Gote/utils"
	"log"
)

var (
	v = validator.New()
)

func Validator() gin.HandlerFunc {
	return func(c *gin.Context) {

		e := en.New()
		uni := ut.New(e, e)

		trans, _ := uni.GetTranslator("en")
		if err := enTranslations.RegisterDefaultTranslations(v, trans); err != nil {
			log.Fatalln("failed to register translator")
		}

		registerValidation("is_admin", func(fl validator.FieldLevel) bool {
			return fl.Field().String() == "admin"
		})

		translateOverride(trans, "required", "{0} must be provided!")
		translateOverride(trans, "is_admin", "user must be admin")

		c.Set(utils.ValidatorKey, v)
		c.Set(utils.TranslatorKey, trans)
	}
}

func translateOverride(trans ut.Translator, validateTag string, msg string) {
	if err := v.RegisterTranslation(validateTag, trans, func(ut ut.Translator) error {
		return ut.Add(validateTag, msg, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(validateTag, fe.Field())
		return t
	}); err != nil {
		log.Fatalln("failed to register trans override")
	}

}

func registerValidation(validateTag string, f validator.Func) {
	if err := v.RegisterValidation(validateTag, f); err != nil {
		log.Fatalln("failed to register validation")
	}
}
