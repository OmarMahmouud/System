package courses

import (
	"github.com/thedevsaddam/govalidator"
	"net/http"
	"net/url"
)

func ValidStore(r *http.Request, request *StoreRequest) (bool, url.Values) {
	rules := govalidator.MapData{
		"course_name": []string{"unique:courses", "required"},
		"description": []string{"required"},
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

func ValidUbdate(r *http.Request, request *UpdatedRequest) (bool, url.Values) {
	rules := govalidator.MapData{
		"course_name": []string{"unique:courses"},
	}
	messages := govalidator.MapData{
		"course_name": []string{"unique: this field must be unique"},
	}
	opts := govalidator.Options{
		Request:         r,
		Messages:        messages,
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
