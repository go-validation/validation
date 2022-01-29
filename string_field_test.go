package validation_test

import (
	"testing"

	"github.com/go-validation/validation"
	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	var i = ""
	errs, err := validation.String("country", &i).NotEmpty().Validate()
	assert.Nil(t, err)
	assert.NotNil(t, errs)

	//assert.Equals()
}
