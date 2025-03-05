package params_validator

import (
	"fmt"
	regexp "github.com/dlclark/regexp2"
	"github.com/go-playground/validator/v10"
	"strconv"
	"strings"
)

const OPTIONAL = "optional"

// baseParam 获取长度校验的最小和最大长度
// baseParam 判断是否需要校验长度
func baseLengthParam(fl validator.FieldLevel) (int, int, bool) {
	param := fl.Param()
	if param == "" {
		return 0, 0, false
	}

	params := strings.Split(param, "-")

	if len(params) != 2 {
		return 0, 0, false
	}

	minNum, err := strconv.Atoi(params[0])
	if err != nil {
		return 0, 0, false
	}
	if minNum == 0 {
		return 0, 0, false
	}

	maxNum, err := strconv.Atoi(params[1])
	if err != nil {
		return 0, 0, false
	}

	if minNum > maxNum {
		return 0, 0, false
	}

	return minNum, maxNum, true
}

func xPhone(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^1\d{10}$`, regexp.None)
	ok, err := re.MatchString(fl.Field().String())
	if err != nil {
		return false
	}
	return ok
}

func xPassword(fl validator.FieldLevel) bool {
	minNum, maxNum, ok := baseLengthParam(fl)
	if ok != true {
		return false
	}
	reg := `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?])[a-zA-Z\d!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]` + fmt.Sprintf(`{%d,%d}$`, minNum, maxNum)
	re := regexp.MustCompile(reg, regexp.None)
	do, errX := re.MatchString(fl.Field().String())
	if errX != nil {
		return false
	}

	return do
}

// xStr 普通字符串
func xStr(fl validator.FieldLevel) bool {
	minNum, maxNum, ok := baseLengthParam(fl)
	if ok != true {
		return false
	}

	reg := `^(?!\s)(?!.*\s$)[\s\S]` + fmt.Sprintf("{%d,%d}", minNum, maxNum) + `$`
	re := regexp.MustCompile(reg, regexp.None)
	do, err := re.MatchString(fl.Field().String())
	if err != nil {
		return false
	}
	return do
}

func xStrWithoutZh(fl validator.FieldLevel) bool {
	minNum, maxNum, ok := baseLengthParam(fl)
	if ok != true {
		return false
	}
	reg := `^(?! )[^\u4e00-\u9fa5]` +
		fmt.Sprintf(`{%d,%d}$`, minNum, maxNum) + `(?<! )$`
	re := regexp.MustCompile(reg, regexp.None)
	do, err := re.MatchString(fl.Field().String())
	if err != nil {
		return false
	}
	return do
}

func xStrWithoutZhAndSpec(fl validator.FieldLevel) bool {
	minNum, maxNum, ok := baseLengthParam(fl)
	if ok != true {
		return false
	}
	reg := `(?! )[A-Za-z0-9 ]` +
		fmt.Sprintf(`{%d,%d}$`, minNum, maxNum) + `(?<! )$`
	re := regexp.MustCompile(reg, regexp.None)
	do, err := re.MatchString(fl.Field().String())
	if err != nil {
		return false
	}
	return do
}

func xStrWithoutSpec(fl validator.FieldLevel) bool {
	minNum, maxNum, ok := baseLengthParam(fl)
	if ok != true {
		return false
	}
	reg := `^(?! )[A-Za-z0-9\u4e00-\u9fa5 ]` + fmt.Sprintf(`{%d,%d}$`, minNum, maxNum) + `(?<! )$`
	re := regexp.MustCompile(reg, regexp.None)
	do, err := re.MatchString(fl.Field().String())
	if err != nil {
		return false
	}
	return do
}

func xStrWithoutSpecAndSpace(fl validator.FieldLevel) bool {
	minNum, maxNum, ok := baseLengthParam(fl)
	if ok != true {
		return false
	}
	reg := `^[A-Za-z0-9\u4e00-\u9fa5]` + fmt.Sprintf(`{%d,%d}$`, minNum, maxNum)
	re := regexp.MustCompile(reg, regexp.None)
	do, err := re.MatchString(fl.Field().String())
	if err != nil {
		return false
	}
	return do
}
