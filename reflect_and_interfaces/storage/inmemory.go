package storage

import "reflect_and_interfaces/user"

type Memory struct {
	users []user.User
}

func (m *Memory) AddUser(u user.User) {
	m.users = append(m.users, u)
}

func (m *Memory) CreateUser(u user.User) {
	m.users = append(m.users, u)
}

func (m *Memory) ListUsers() []user.User {

}

func (m *Memory) GetUserByID(id uint) user.User {
	for _, user := range m.users {
		if user.ID == id {
			return user
		}
	}

	return user.User{}
}
