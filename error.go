package validation

type Error interface {
	Property() string
	Value() interface{}
	Violations() []Violation
	Children() Errors
}

type Errors []Error

type Violation struct {
	MsgKey string
	Params map[string]interface{}
}

type fieldError struct {
	fieldName  string
	value      interface{}
	violations []Violation
	children   Errors
}

func newFieldError(key string, value interface{}, code string, params map[string]interface{}) *fieldError {
	violations := make([]Violation, 1)
	violations[0] = Violation{MsgKey: code, Params: params}
	return &fieldError{fieldName: key, value: value, violations: violations}
}

func (fe *fieldError) Property() string {
	return fe.fieldName
}

func (fe *fieldError) Value() interface{} {
	return fe.value
}

func (fe *fieldError) Violations() []Violation {
	return fe.violations
}

func (fe *fieldError) Children() Errors {
	return fe.children
}
