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
	saver := &FooSaver{
		Validator: validator,
		Database: database,
	}

	err := saver.SaveFoo(foo)
	if err != nil {
		log.Fatal(err)
	}
}

type FooSaver struct {
	Validator validation.Validator
	Database db.Database
}

// error handling omitted for brevity
func (f *FooSaver) SaveFoo(foo Foo) error {

	f.Validator.Validate(foo)

	query := `INSERT INTO foos (id, name) VALUES (?, ?)`
	f.Database.Insert(query, []string{foo.ID, foo.Name})

	return nil
}
