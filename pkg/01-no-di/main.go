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
	err := SaveFoo(foo)
	if err != nil {
		log.Fatal(err)
	}
}

// error handling omitted for brevity
func SaveFoo(foo Foo) error {

	validator := validation.GetValidator()
	validator.Validate(foo)

	query := `INSERT INTO foos (id, name) VALUES (?, ?)`
	database := db.GetDatabase()
	database.Insert(query, []string{foo.ID, foo.Name})

	return nil
}
