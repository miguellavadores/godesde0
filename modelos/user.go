package modelos

import (
	"time"
)

type User struct {
	Id int
	Name string
	CreatedAt time.Time
	status bool
}

func (usuario *User) AddUser(id int, name string, createdAt time.Time, status bool) {
	usuario.Id = id
	usuario.Name = name
	usuario.CreatedAt = createdAt
	usuario.status = status
}