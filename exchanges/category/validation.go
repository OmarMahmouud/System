package category

import (
	"github.com/thedevsaddam/govalidator"
	"net/http"
	"net/url"
)

func ValidStore(r *http.Request, request *StoreCategory) (bool, url.Values) {
	rules := govalidator.MapData{
		"name": []string{"unique:categories", "required"},
	}
	opts := govalidator.Options{
		Request:         r,
		Data:            request,
		Rules:           rules,
		RequiredDefault: true, //force user to fill all the inputs
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
		"name": []string{"unique:categories"},
	}
	messages := govalidator.MapData{
		"name": []string{"unique: this field must be unique"},
	}
	opts := govalidator.Options{
		Request:         r,
		Messages:        messages,
		Data:            request,
		Rules:           rules,
		RequiredDefault: true, //force user to fill all the inputs
	}
	v := govalidator.New(opts)
	e := v.ValidateJSON()
	if len(e) > 0 {
		return false, e
	}
	return true, e
}
