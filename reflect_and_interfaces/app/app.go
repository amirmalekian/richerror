package app

import (
	"fmt"
	"reflect_and_interfaces/user"
)

type UserStore interface {
	CreateUser(u user.User)
	ListUsers() []user.User
	GetUserByID(id uint) user.User
}

type TodoStore interface {
	CreateTodo()
	GetTodoByID()
}

type App struct {
	Name string
	// StorageFilePath string
	// InMemoryStorage storage.Memory
	UserStorage UserStore
}

func (a App) CreateUser(u user.User) {
	if u.Name == "" {
		fmt.Println("name can't be empty")

		return
	}
	// var fileHandler *os.File
	// if f, err := os.OpenFile(a.StorageFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777); err != nil {
	// 	fmt.Println("can't open file", err)

	// 	return
	// } else {
	// 	fileHandler = f
	// }

	// defer fileHandler.Close()

	// serialize user
	// data, mErr := json.Marshal(u)
	// if mErr != nil {
	// 	fmt.Println("can't marshal user data", mErr)

	// 	return
	// }

	// a.InMemoryStorage.AddUser(u)

	a.UserStorage.CreateUser(u)

	// if _, wErr := fileHandler.Write(data); wErr != nil {
	// 	fmt.Println("can't write to the file", mErr)

	// 	return
	// }
}

// func (a App) ListUsers() []User {

// }

// func (a App) GetUserByID(id uint) User {

// }
