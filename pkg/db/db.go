package db

import "fmt"

type Database interface {
	Insert(query string, args []string) error
}

func GetDatabase() Database {
	return &database{}
}

type database struct {}

func (database) Insert(query string, args []string) error {
	// fake!
	fmt.Printf("Inserting... %s %v\n", query, args)
	return nil
}

