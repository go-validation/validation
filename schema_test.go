package validation_test

import (
	"fmt"
	"testing"

	"github.com/go-validation/validation"
	"github.com/stretchr/testify/assert"
)

type Address struct {
	Country, Province, City string
}

func AddressSchema(a *Address) validation.Schema {
	return validation.Schema{
		validation.String("country", &a.Country).NotBlank(),
		validation.String("city", &a.Country).NotBlank(),
	}
}

type Person struct {
	Name    string
	Age     int
	Address Address
}

func (p *Person) Schema() validation.Schema {
	return validation.Schema{
		validation.String("name", &p.Name).NotBlank(),
		validation.Num("age", &p.Age).NotZero(),
		validation.StructBy("address", &p.Address, AddressSchema),
	}
}

func TestSchema(t *testing.T) {
	var person = Person{
		Name:    "John",
		Age:     0,
		Address: Address{},
	}
	vErrs, err := (&person).Schema().Validate()

	assert.Nil(t, err)
	assert.NotNil(t, vErrs)

	fmt.Printf("%+v", vErrs)

}
