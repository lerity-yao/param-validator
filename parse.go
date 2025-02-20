package params_validator

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"strings"
)

// HttpxParseValidator 构建 httpx Validator
type HttpxParseValidator struct {
	Validator
}

// MustNewHttpxParseValidator 构建 httpx Validator 实例，捕捉错误
func MustNewHttpxParseValidator(conf Conf) *HttpxParseValidator {
	vd := MustNewValidate(conf)
	h := &HttpxParseValidator{*vd}
	err := h.InitRegisterValidation()
	logx.Must(err)
	return h
}

// Validate 构建 httpx Validator的 Validate
func (v *HttpxParseValidator) Validate(r *http.Request, data any) (err error) {
	defer func() {
		if r := recover(); r != nil {
			// 捕获 panic 并将其转换为 error
			err = fmt.Errorf("validation panic: %v", r)
		}
	}()
	err = v.Validator.Validator.Struct(data)
	if err != nil {
		var errMsg []string
		for _, e := range err.(validator.ValidationErrors) {
			errMsg = append(errMsg, e.Translate(v.Validator.Trans))
		}
		return errors.New(strings.Join(errMsg, ","))
	}

	return nil
}

func (v *HttpxParseValidator) InitRegisterValidation() (err error) {
	if err = v.xPhone(); err != nil {
		return err
	}

	if err = v.xPassword(); err != nil {
		return err
	}

	if err = v.xStr(); err != nil {
		return err
	}

	if err = v.xStrWithoutSpec(); err != nil {
		return err
	}

	if err = v.xStrWithoutZh(); err != nil {
		return err
	}

	if err = v.xStrWithoutSpecAndSpace(); err != nil {
		return err
	}

	if err = v.xStrWithoutZhAndSpec(); err != nil {
		return err
	}
	return nil
}

// xPhone 注册自定义 xPhone 方法
// xPhone 为手机号校验规则, 校验其是否为1开头，长度为11位的数字字符串
func (v *HttpxParseValidator) xPhone() error {
	if err := v.Validator.RegisterValidation("xPhone", xPhone); err != nil {
		return err
	}
	if err := v.Validator.RegisterTranslation("xPhone", "{0}必须为手机号，1开头，长度为11位", false); err != nil {
		return err
	}

	return nil
}

// xPassword 注册自定义 xPassword 方法
// xPassword 为密码校验规则, 使用方法为 validate:"xPassword=8-15" 代表 字符串长度为 8到15位。左右都是闭区间
// xPassword 为密码校验规则, 校验其长度位，需由字母（同时要大小和写）、数字、特殊字符串三种组成，不能使用空格、中文
func (v *HttpxParseValidator) xPassword() error {
	if err := v.Validator.RegisterValidation("xPassword", xPassword); err != nil {
		return err
	}
	if err := v.Validator.RegisterTranslation("xPassword", "{0}长度{1}，需由字母（区分大小写）、数字、特殊字符串三种组成，不能使用空格、中文", false); err != nil {
		return err
	}
	return nil
}

// xStr 注册自定义 xStr 方法
// xStr 为字符串校验方法，使用方法为 validate:"xStr=1-300" 代表字符串长度为 1-300位，左右都为闭区间
// xStr 长度自定义
// xStr 首尾不能有空格，中间可以有空格
// xStr 允许中文，特殊字符串，英文，数字
func (v *HttpxParseValidator) xStr() error {
	if err := v.Validator.RegisterValidation("xStr", xStr); err != nil {
		return err
	}
	if err := v.Validator.RegisterTranslation("xStr", "{0}长度{1}，首尾不能有空格", false); err != nil {
		return err
	}
	return nil
}

// xStrWithoutZh 注册自定义 xStrWithoutZh 方法
// xStrWithoutZh 为字符串校验方法，使用方法为 validate:"xStrWithoutZh=1-300" 代表字符串长度为 1-300位，左右都为闭区间
// xStrWithoutZh 长度自定义
// xStrWithoutZh 首尾不能有空格，中间可以有空格
// xStrWithoutZh 不能包含中文
// xStrWithoutZh 允许特殊字符串，英文，数字
func (v *HttpxParseValidator) xStrWithoutZh() error {
	if err := v.Validator.RegisterValidation("xStrWithoutZh", xStrWithoutZh); err != nil {
		return err
	}
	if err := v.Validator.RegisterTranslation("xStrWithoutZh", "{0}长度{1}，首尾不能有空格，不能包含中文", false); err != nil {
		return err
	}
	return nil
}

// xStrWithoutZhAndSpec
// xStrWithoutZhAndSpec 注册自定义 xStrWithoutZhAndSpec 方法
// xStrWithoutZhAndSpec 为字符串校验方法，使用方法为 validate:"xStrWithoutZhAndSpec=1-300" 代表字符串长度为 1-300位，左右都为闭区间
// xStrWithoutZhAndSpec 长度自定义
// xStrWithoutZhAndSpec 首尾不能有空格，中间可以有空格
// xStrWithoutZhAndSpec 不能包含中文
// xStrWithoutZhAndSpec 不能包含特殊字符 !@#$%^&*()_+\-=[]{};':"\\|,.<>/?等
// xStrWithoutZhAndSpec 允许字母和数字
func (v *HttpxParseValidator) xStrWithoutZhAndSpec() error {
	if err := v.Validator.RegisterValidation("xStrWithoutZhAndSpec", xStrWithoutZhAndSpec); err != nil {
		return err
	}
	if err := v.Validator.RegisterTranslation("xStrWithoutZhAndSpec", "{0}长度{1}，首尾不能有空格，不能包含中文和特殊字符串", false); err != nil {
		return err
	}
	return nil
}

// xStrWithoutSpec 注册自定义 xStrWithoutSpec 方法
// xStrWithoutSpec 为字符串校验方法，使用方法为 validate:"xStrWithoutSpec=1-300" 代表字符串长度为 1-300位，左右都为闭区间
// xStrWithoutSpec 长度自定义
// xStrWithoutSpec 首尾不能有空格，中间可以有空格
// xStrWithoutSpec 不能包含特殊字符 !@#$%^&*()_+\-=[]{};':"\\|,.<>/?等
// xStrWithoutSpec 可以包含其他字符（如字母、数字、空格、中文等）。
func (v *HttpxParseValidator) xStrWithoutSpec() error {
	if err := v.Validator.RegisterValidation("xStrWithoutSpec", xStrWithoutSpec); err != nil {
		return err
	}
	if err := v.Validator.RegisterTranslation("xStrWithoutSpec", "{0}长度{1}，首尾不能有空格，不能包含特殊字符串", false); err != nil {
		return err
	}
	return nil
}

// xStrWithoutSpecAndSpace 注册自定义 xStrWithoutSpecAndSpace 方法
// xStrWithoutSpecAndSpace 为字符串校验方法，使用方法为 validate:"xStrWithoutSpecAndSpace=1-300" 代表字符串长度为 1-300位，左右都为闭区间
// xStrWithoutSpecAndSpace 长度自定义
// xStrWithoutSpecAndSpace 不能有空格
// xStrWithoutSpecAndSpace 不能包含特殊字符 !@#$%^&*()_+\-=[]{};':"\\|,.<>/?等
// xStrWithoutSpecAndSpace 可以包含其他字符（如字母、数字、中文等）。
func (v *HttpxParseValidator) xStrWithoutSpecAndSpace() error {
	if err := v.Validator.RegisterValidation("xStrWithoutSpecAndSpace", xStrWithoutSpecAndSpace); err != nil {
		return err
	}
	if err := v.Validator.RegisterTranslation("xStrWithoutSpecAndSpace", "{0}长度{1}，不能有空格，不能包含特殊字符串", false); err != nil {
		return err
	}
	return nil
}
