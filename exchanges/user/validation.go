package user

import (
	"github.com/thedevsaddam/govalidator"
	"net/http"
	"net/url"
)

func ValidStore(r *http.Request, request *SignupRequest) (bool, url.Values) {
	rules := govalidator.MapData{
		"email":     []string{"unique:users", "required", "email"},
		"password":  []string{"required"},
		"user_name": []string{"required"},
	}
	opts := govalidator.Options{
		Request:         r,
		Data:            request,
		Rules:           rules,
		RequiredDefault: true,
	}
	v := govalidator.New(opts)
	e := v.ValidateJSON()
	if len(e) > 0 {
		return false, e
	}
	return true, e
}
