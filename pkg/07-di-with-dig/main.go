package main

import (
	"github.com/dexhorthy/golang-di/pkg/db"
	"github.com/replicatedcom/replicated/operator/log"
	"github.com/dexhorthy/golang-di/pkg/validation"
	"go.uber.org/dig"
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

	injector := dig.New()
	injector.Provide(db.GetDatabase())
	injector.Provide(validation.GetValidator())
	injector.Provide(NewFooSaver)
	injector.Provide(NewBarSaver)
	injector.Provide(NewApp)

	err := injector.Invoke(func(app *App) error {
		return app.Save(foo, Bar{ID: "bar", Name: "my bar"})
	})

	if err != nil {
		log.Fatal(err)
	}
}

func NewFooSaver(validator validation.Validator, database db.Database) *FooSaver {
	return &FooSaver{
		Validator: validator,
		Database: database,
	}
}

func NewBarSaver(database db.Database) (*BarSaver, error) {
	return &BarSaver{
		Database: database,
	}, nil // but it could fail
}

func NewApp(fooSaver *FooSaver, barSaver *BarSaver, database db.Database) *App {
	return &App{
		FooSaver: fooSaver,
		BarSaver: barSaver,
		Database: database,
	}
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

