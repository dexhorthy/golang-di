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

type Bar struct {
	ID   string
	Name string
}

func main() {
	foo := Foo{
		ID: "12345",
		Name: "my foo",
	}

	app := GetApp()

	err := app.Save(foo, Bar{ID: "bar", Name: "my bar"})
	if err != nil {
		log.Fatal(err)
	}
}

func GetApp() *App {
	// bootstrap dependencies
	validator := validation.GetValidator()
	database := db.GetDatabase()
	fooSaver := &FooSaver{
		Validator: validator,
		Database: database,
	}
	barSaver := &BarSaver{
		Database: database,
	}

	app := &App{
		FooSaver: fooSaver,
		BarSaver: barSaver,
		Database: database,
	}

	return app
}

type App struct {
	FooSaver *FooSaver
	BarSaver *BarSaver
	Database db.Database
}

type BarSaver struct {
	Database db.Database
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

func (a *App) Save(foo Foo, bar Bar) error {
	// ... do stuff with all the fields
}

