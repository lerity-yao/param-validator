package params_validator

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"github.com/zeromicro/go-zero/core/logx"
	"reflect"
	"strings"
)

type (
	Validator struct {
		Validator *validator.Validate
		Trans     ut.Translator
	}
)

// MustNewValidate 构建 validator，捕捉错误
func MustNewValidate(conf Conf) *Validator {
	v, err := NewValidator(conf)
	logx.Must(err)
	return v
}

// NewValidator 新构建 NewValidator
func NewValidator(conf Conf) (*Validator, error) {
	vd := validator.New()
	// 注册一个函数，获取struct tag里自定义的label作为字段名
	vd.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("form"), ",", 2)[0]
		if name == "-" || name == "" {
			name = strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" || name == "" {
				name = strings.SplitN(fld.Tag.Get("header"), ",", 2)[0]
				if name == "-" || name == "" {
					name = strings.SplitN(fld.Tag.Get("path"), ",", 2)[0]
					if name == "-" || name == "" {
						return ""
					}
				}
			}
		}
		return name
	})

	v := &Validator{Validator: vd}

	if conf.ZhTrans {
		zhCh := zh.New()
		uni := ut.New(zhCh)
		trans, _ := uni.GetTranslator("zh") // 验证器注册翻译器
		err := zhTrans.RegisterDefaultTranslations(vd, trans)
		if err != nil {
			return nil, err
		}
		v.Trans = trans
	}

	if !conf.ZhTrans {
		enCh := en.New()
		uni := ut.New(enCh)
		trans, _ := uni.GetTranslator("en") // 验证器注册翻译器
		err := zhTrans.RegisterDefaultTranslations(vd, trans)
		if err != nil {
			return nil, err
		}
		v.Trans = trans
	}

	return v, nil
}

// RegisterValidation 暴露注册自定义 validation
func (v *Validator) RegisterValidation(tag string, fn validator.Func, callValidationEvenIfNull ...bool) error {
	err := v.Validator.RegisterValidation(tag, fn, callValidationEvenIfNull...)
	logx.Must(err)
	return err
}

// translate 自定义翻译器
func (v *Validator) translate(trans ut.Translator, fe validator.FieldError) string {
	msg, err := trans.T(fe.Tag(), fe.Field(), fe.Param())
	if err != nil {
		return fe.(error).Error()
	}
	return msg
}

// registerTranslator 注意自定义翻译tag
func (v *Validator) registerTranslator(tag, msg string, override bool) validator.RegisterTranslationsFunc {
	return func(trans ut.Translator) error {
		if err := trans.Add(tag, msg, override); err != nil {
			return err
		}
		return nil
	}
}

// RegisterTranslation 暴露注册自定义 validation自定义返回错误结果
func (v *Validator) RegisterTranslation(tag, msg string, override bool) error {
	if v.Trans == nil {

	}
	err := v.Validator.RegisterTranslation(tag, v.Trans, v.registerTranslator(tag, msg, override), v.translate)
	return err
}
