package newinmemory

import "reflect_and_interfaces/user"

type Store struct {
	users map[uint]user.User
}

func (s *Store) CreateUser(u user.User) {
}

func (s *Store) ListUsers() []user.User {

}

func (s *Store) GetUserByID(id uint) user.User {
	return s.users[id]
}
