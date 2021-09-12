package validation

import "strings"

type StringField[T StringType] struct {
	FieldValidator[T]
}
type StringType interface {
	~*string
}

var _ Validatable = &StringField[*string]{}

func notEmptyRule[T StringType]() *rule[T] {
	return &rule[T]{
		code: MsgKeyStringNotEmpty,
		validatorFn: func(value T) (bool, error) {
			return value != nil && len(*value) > 0, nil
		},
	}
}

func notBlankRule[T StringType]() *rule[T] {
	return &rule[T]{
		code: MsgKeyStringNotEmpty,
		validatorFn: func(value T) (bool, error) {
			return value != nil && len(strings.TrimSpace(*value)) > 0, nil
		},
	}
}

func minLengthRule[T StringType](minLength int) *rule[T] {
	return &rule[T]{
		code: MsgKeyStringMinLength,
		validatorFn: func(value T) (bool, error) {
			return len(*value) >= minLength, nil
		},
	}
}

func maxLengthRule[T StringType](max int) *rule[T] {
	return &rule[T]{
		code: MsgKeyStringMaxLength,
		validatorFn: func(value T) (bool, error) {
			return len(*value) <= max, nil
		},
	}
}

func lengthBetweenRule[T StringType](min, max int) *rule[T] {
	var params = make(map[string]interface{})
	params["min"] = min
	params["max"] = max
	return &rule[T]{
		code: MsgKeyStringMaxLength,
		validatorFn: func(value T) (bool, error) {
			return len(*value) >= min && len(*value) <= max, nil
		},
		params: params,
	}
}

func String[T StringType](fieldName string, value T) *StringField[T] {
	return &StringField[T]{
		FieldValidator: FieldValidator[T]{
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
