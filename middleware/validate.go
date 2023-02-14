package middleware

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"regexp"
)

func ValidateFunction() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("passwordAllow", passwordAllow)
		if err != nil {
			return
		}
		err = v.RegisterValidation("usernameAllow", usernameAllow)
		if err != nil {
			return
		}
	}
}

var passwordAllow validator.Func = func(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	pattern := "[\\w]{8,}"
	re, _ := regexp.MatchString(pattern, password)
	return re
}

var usernameAllow validator.Func = func(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	pattern := "(\\w+)$"
	re, _ := regexp.MatchString(pattern, password)
	return re
}
