package main

import (
	"reflect_and_interfaces/app"
	newinmemory "reflect_and_interfaces/new_in_memory"
	"reflect_and_interfaces/user"
)

type name interface {
}

func main() {
	// var a = &storage.Memory{}

	var newMemoryStorage = &newinmemory.Store{}

	application := app.App{
		Name: "sample-app",
		// StorageFilePath: "./data.txt",
		// InMemoryStorage: storage.Memory{},
		UserStorage: newMemoryStorage,
	}

	application.CreateUser(user.User{
		ID:   1,
		Name: "Amirhossein",
	})
}
