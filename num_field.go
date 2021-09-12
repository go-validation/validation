package validation

type NumField[T NumValue] struct {
	FieldValidator[T]
}

type NumValue interface {
	~*int | ~*int8 | ~*int16 | ~*int32 | ~*int64 |
		~*uint | ~*uint8 | ~*uint16 | ~*uint32 | ~*uint64 |
		~*float32 | ~*float64
}

func Num[T NumValue](fieldName string, value T) *NumField[T] {
	return &NumField[T]{
		FieldValidator[T]{
			fieldName: fieldName,
			value:     value,
		},
	}
}

func NotZeroRule[T NumValue]() ValidationRule[T] {
	return &rule[T]{
		code: MsgKeyNumberNotZero,
		validatorFn: func(value T) (bool, error) {
			var zero = *new(T)
			return value != nil && value != zero, nil
		},
	}
}

func (f *NumField[T]) NotZero() *NumField[T] {
	f.rules = append(f.rules, NotZeroRule[T]())
	return f
}
