package http

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTrans "github.com/go-playground/validator/v10/translations/en"
)

type ValidationError struct {
	errors map[string]string
}

func (e ValidationError) Error() string {
	return "The given data was invalid."
}

func (e ValidationError) Errors() map[string]string {
	return e.errors
}

type Validator struct {
	validate *validator.Validate
	trans    *ut.Translator
}

var (
	v    *Validator
	once sync.Once
)

// NewValidator creates a new instance of the Validator struct
func NewValidator(r *http.Request, params any) error {
	once.Do(func() {
		locale := en.New()
		uni := ut.New(locale, locale)
		validate := validator.New(validator.WithRequiredStructEnabled())
		trans, ok := uni.GetTranslator("en")
		if ok {
			_ = enTrans.RegisterDefaultTranslations(validate, trans)
		}
		v = &Validator{validate, &trans}
	})
	return v.exec(r, params)
}

// exec is a method of the Validator struct that validates the given request and parameters
func (v Validator) exec(r *http.Request, params any) error {
	if err := json.NewDecoder(r.Body).Decode(params); err != nil {
		return err
	}
	if err := v.validate.Struct(params); err != nil {
		vErr := ValidationError{
			errors: err.(validator.ValidationErrors).Translate(*v.trans),
		}
		return vErr
	}
	return nil
}
