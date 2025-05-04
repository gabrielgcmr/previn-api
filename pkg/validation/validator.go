package validation

import (
	"github.com/go-playground/locales/pt_BR"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	pt_translations "github.com/go-playground/validator/v10/translations/pt_BR"
)

var (
	Validate   *validator.Validate
	Translator ut.Translator
)

func Init() error {
	Validate = validator.New()

	// Configurar tradutor
	locale := pt_BR.New()
	uni := ut.New(locale, locale)
	trans, _ := uni.GetTranslator("pt_BR")
	Translator = trans

	// Registrar traduções
	return pt_translations.RegisterDefaultTranslations(Validate, Translator)
}

// TranslateErrors translates validation errors into a map of field names and error messages.
func TranslateErrors(err error) map[string]string {
	errs := make(map[string]string)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			errs[fieldError.Field()] = fieldError.Error()
		}
	}
	return errs
}
