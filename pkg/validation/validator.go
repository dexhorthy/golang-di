package validation

import "fmt"


type Validator interface {
	Validate(obj interface{}) error
}

func GetValidator() Validator {
	return &validator{}
}

type validator struct {}

func (validator) Validate(obj interface{}) error {
	// fake!
	fmt.Printf("Validating %s", obj)
	return nil
}
