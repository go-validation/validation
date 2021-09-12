package validation_test

import (
	"testing"

	"github.com/go-validation/validation"
	"github.com/stretchr/testify/assert"
)

var _ validation.Validatable = &validation.NumField[*int]{}
var _ validation.Validatable = &validation.NumField[*int32]{}
var _ validation.Validatable = &validation.NumField[*int64]{}
var _ validation.Validatable = &validation.NumField[*int16]{}
var _ validation.Validatable = &validation.NumField[*int8]{}

func TestNumber(t *testing.T) {
	var i = 0
	errs, err := validation.Num("country", &i).NotZero().Validate()

	assert.Nil(t, err)
	assert.NotNil(t, errs)

}
