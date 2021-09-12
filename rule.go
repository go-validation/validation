package validation

type ValidationRule[T any] interface {
	MsgKey() string
	IsValid(value T) (bool, error)
	Params() map[string]interface{}
	WithMsgKey(string) ValidationRule[T]
}

func Rule[T any](code string, validatorFn func(T) (bool, error), params map[string]interface{}) ValidationRule[T] {
	return &rule[T]{
		code:        code,
		params:      params,
		validatorFn: validatorFn,
	}
}

type rule[T any] struct {
	code        string
	params      map[string]interface{}
	validatorFn func(value T) (bool, error)
}

func (r *rule[T]) MsgKey() string {
	return r.code
}

func (r *rule[T]) WithMsgKey(key string) ValidationRule[T] {
	return &rule[T]{
		code:        key,
		params:      r.params,
		validatorFn: r.validatorFn,
	}
}

func (r *rule[T]) Params() map[string]interface{} {
	return r.params
}

func (r *rule[T]) IsValid(value T) (bool, error) {
	return r.validatorFn(value)
}
