package repository

import (
	models "CardozoCasariegoLuciano/StudyNotes/Models"
	"fmt"
	"sync"
)

var db *memory
var once sync.Once

type memory struct {
	users []models.User
}

func NewMemory() *memory {
	once.Do(func() {
		fmt.Println("Pasa por aca memory")
		db = &memory{users: []models.User{}}
	})
	return db
}

func (memory *memory) Save(user models.User) models.User {
	user.Id = len(memory.users) + 1
	memory.users = append(memory.users, user)
	return user
}

func (memory *memory) ListAll() []models.User {
	return memory.users
}

func (memory *memory) FindUserByEmail(email string) models.User {
	var userFinded models.User
	for _, user := range memory.users {
		if user.Email == email {
			userFinded = user
			break
		}
	}

	return userFinded
}
