package validation_test

import (
	"github.com/go-validation/validation"
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
