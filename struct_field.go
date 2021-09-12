package validation

type StructField[T any] struct {
	FieldValidator[T]
}

func Struct[T SchemaI](fieldName string, val T) Validatable {
	return StructBy(fieldName, val, SchemaFn[T](func(t T) Schema {
		return t.Schema()
	}))
}

func StructBy[T any](fieldName string, val T, s SchemaFn[T]) Validatable {
	return ValidatableFunc(
		func() (Error, error) {
			errs, err := s(val).Validate()

			if err != nil {
				return nil, err
			}
			if errs != nil {
				fe := fieldError{
					fieldName: fieldName,
					value:     val,
					children:  errs,
				}
				return &fe, nil
			}
			return nil, nil
		})
}
