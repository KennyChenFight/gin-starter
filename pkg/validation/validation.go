package validation

import (
	"errors"
	"reflect"
	"strings"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh_Hant"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTWTranslations "github.com/go-playground/validator/v10/translations/zh_tw"
)

type registerLocaleTranslation func(v *validator.Validate, trans ut.Translator) (err error)

type localeTranslation struct {
	locale              locales.Translator
	translationRegister registerLocaleTranslation
}

type Locale string

const (
	En     Locale = "en"
	ZhHant Locale = "zh_Hant"
)

var supportedLocales = map[Locale]localeTranslation{
	En:     {locale: en.New(), translationRegister: enTranslations.RegisterDefaultTranslations},
	ZhHant: {locale: zh_Hant.New(), translationRegister: zhTWTranslations.RegisterDefaultTranslations},
}

func NewValidationTranslator(ginBindingValidator *validator.Validate, defaultLanguage string) (*ValidationTranslator, error) {
	if _, ok := supportedLocales[Locale(defaultLanguage)]; !ok {
		return nil, errors.New("not supported language")
	}

	var localeTranslators []locales.Translator
	for _, localeTranslation := range supportedLocales {
		localeTranslators = append(localeTranslators, localeTranslation.locale)
	}
	uni := ut.New(supportedLocales[Locale(defaultLanguage)].locale, localeTranslators...)

	for locale, localeTranslation := range supportedLocales {
		trans, found := uni.GetTranslator(string(locale))
		if found {
			err := localeTranslation.translationRegister(ginBindingValidator, trans)
			if err != nil {
				return nil, err
			}
		}
	}

	ginBindingValidator.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	return &ValidationTranslator{universalTranslator: uni, Validate: ginBindingValidator}, nil
}

type ValidationTranslator struct {
	universalTranslator *ut.UniversalTranslator
	Validate            *validator.Validate
}

func (v *ValidationTranslator) Translate(language string, err error) (validator.ValidationErrorsTranslations, error) {
	if language == "" {
		language = "en"
	}
	validationErr, ok := err.(validator.ValidationErrors)
	if !ok && language != "" {
		return nil, nil
	}
	trans, found := v.universalTranslator.GetTranslator(language)
	if !found {
		return nil, errors.New("not supported language")
	}

	return validationErr.Translate(trans), nil
}
