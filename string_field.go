package validation

import (
	"fmt"
	"regexp"
	"strings"
)

type StringField[T StringType] struct {
	FieldValidator[*T]
}
type StringType interface {
	~string
}

var _ Validatable = &StringField[string]{}

const (
	emailPattern          string = "^(((([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|((\\x22)((((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(([\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(\\([\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(\\x22)))@((([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
	creditCardPattern     string = "^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|(222[1-9]|22[3-9][0-9]|2[3-6][0-9]{2}|27[01][0-9]|2720)[0-9]{12}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\\d{3})\\d{11}|6[27][0-9]{14})$"
	alphaPattern          string = "^[a-zA-Z]+$"
	alphaNumericPattern   string = "^[a-zA-Z0-9]+$"
	numericPattern        string = "^[0-9]+$"
	allLowerCasePattern   string = ".*[[:lower:]]"
	allUpperCasePattern   string = ".*[[:upper:]]"
	lowerSnakeCasePattern string = "^[a-z1-9_]+$"
	upperSnakeCasePattern string = "^[A-Z1-9_]+$"
)

var (
	emailRegEx        = regexp.MustCompile(emailPattern)
	creditCardRegEx   = regexp.MustCompile(creditCardPattern)
	alphaRegEx        = regexp.MustCompile(alphaPattern)
	alphaNumericRegEx = regexp.MustCompile(alphaNumericPattern)
	numericRegEx      = regexp.MustCompile(numericPattern)
	allLowerCaseRegEx = regexp.MustCompile(allLowerCasePattern)
	allUpperCaseRegEx = regexp.MustCompile(allUpperCasePattern)
)

func notEmptyRule[T StringType]() ValidationRule[*T] {
	return &rule[*T]{
		code: MsgKeyStringNotEmpty,
		validatorFn: func(value *T) (bool, error) {
			//FIXME: hack for generics bug
			if value == nil {
				return false, nil
			}
			var str = fmt.Sprintf("%v", *value)

			return value != nil && len(str) > 0, nil
		},
	}
}

func notBlankRule[T StringType]() ValidationRule[*T] {
	return &rule[*T]{
		code: MsgKeyStringNotEmpty,
		validatorFn: func(value *T) (bool, error) {
			//FIXME: hack for generics bug
			if value == nil {
				return false, nil
			}
			var str = fmt.Sprintf("%v", *value)

			return len(strings.TrimSpace(string(str))) > 0, nil
		},
	}
}

func minLengthRule[T StringType](minLength int) ValidationRule[*T] {
	return &rule[*T]{
		code: MsgKeyStringMinLength,
		validatorFn: func(value *T) (bool, error) {
			//FIXME: hack for generics bug
			if value == nil {
				return 0 >= minLength, nil
			}
			var str = fmt.Sprintf("%v", *value)

			return len(str) >= minLength, nil
		},
	}
}

func maxLengthRule[T StringType](max int) ValidationRule[*T] {
	return &rule[*T]{
		code: MsgKeyStringMaxLength,
		validatorFn: func(value *T) (bool, error) {
			//FIXME: hack for generics bug
			if value == nil {
				return 0 <= max, nil
			}
			var str = fmt.Sprintf("%v", *value)

			return len(str) <= max, nil
		},
	}
}

func lengthBetweenRule[T StringType](min, max int) ValidationRule[*T] {
	var params = make(map[string]interface{})
	params["Min"] = min
	params["Max"] = max
	return &rule[*T]{
		code: MsgKeyStringMaxLength,
		validatorFn: func(value *T) (bool, error) {
			//FIXME: hack for generics bug
			if value == nil {
				return false, nil
			}
			var str = fmt.Sprintf("%v", *value)

			return len(str) >= min && len(str) <= max, nil
		},
		params: params,
	}
}

func String[T StringType](fieldName string, value *T) *StringField[T] {
	return &StringField[T]{
		FieldValidator: FieldValidator[*T]{
			fieldName: fieldName,
			value:     value,
		},
	}
}

func (f *StringField[T]) NotEmpty() *StringField[T] {
	f.rules = append(f.rules, notEmptyRule[T]())
	return f
}

func (f *StringField[T]) NotBlank() *StringField[T] {
	f.rules = append(f.rules, notBlankRule[T]())
	return f
}

func (f *StringField[T]) MinLength(len int) *StringField[T] {
	f.rules = append(f.rules, minLengthRule[T](len))
	return f
}

func (f *StringField[T]) MaxLength(len int) *StringField[T] {
	f.rules = append(f.rules, maxLengthRule[T](len))
	return f
}

func (f *StringField[T]) LengthBetween(min, max int) *StringField[T] {
	f.rules = append(f.rules, lengthBetweenRule[T](min, max))
	return f
}
