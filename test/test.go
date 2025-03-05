package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	params_validator "github.com/lerity-yao/param-validator"
)

func main() {
	type S struct {
		Name  string `json:"name" validate:"xStrWithoutZh=10-10"`
		Name1 string `json:"name1" validate:"xStrWithoutZhAndSpec=1-10"`
		Name2 string `json:"name2" validate:"xStrWithoutSpec=1-10"`
		Name3 string `json:"name3" validate:"xStrWithoutSpecAndSpace=1-10"`
		Name4 string `json:"name4" validate:"xStr=1-10"`
	}

	s := S{Name: "1234567890", Name1: "z", Name2: "中国和中国和中国和", Name3: "te是st", Name4: "test"}

	vd := params_validator.MustNewHttpxParseValidator(params_validator.Conf{ZhTrans: true})

	err := vd.Validator.Validator.Struct(&s)
	if err != nil {
		var errMsg []string
		for _, e := range err.(validator.ValidationErrors) {
			errMsg = append(errMsg, e.Translate(vd.Validator.Trans))
		}
		fmt.Println(errMsg)
	}

}
