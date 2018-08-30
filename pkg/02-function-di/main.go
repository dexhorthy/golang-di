package main

import (
	"github.com/dexhorthy/golang-di/pkg/db"
	"github.com/replicatedcom/replicated/operator/log"
	"github.com/dexhorthy/golang-di/pkg/validation"
)

type Foo struct {
	ID   string
	Name string
}

func main() {
	foo := Foo{
		ID: "12345",
		Name: "my foo",
	}

	// bootstrap dependencies
	validator := validation.GetValidator()
	database := db.GetDatabase()

	err := SaveFoo(foo, validator, database)
	if err != nil {
		log.Fatal(err)
	}
}

// error handling omitted for brevity
func SaveFoo(foo Foo, validator validation.Validator, database db.Database) error {

	validator.Validate(foo)

	query := `INSERT INTO foos (id, name) VALUES (?, ?)`
	database.Insert(query, []string{foo.ID, foo.Name})

	return nil
}
