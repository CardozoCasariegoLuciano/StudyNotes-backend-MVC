package memory

import (
	models "CardozoCasariegoLuciano/StudyNotes/Models"
	"fmt"
	"sync"
)

var DB_Memory *memory
var once_memory sync.Once

type memory struct {
	users []models.User
}

func NewMemory() *memory {
	once_memory.Do(func() {
		fmt.Println("Pasa por aca memory")
		DB_Memory = &memory{users: []models.User{}}
	})
	return DB_Memory
}

func (memory *memory) SaveUser(user *models.User) error {
	user.CommonModelFields.ID = uint(len(memory.users) + 1)
	memory.users = append(memory.users, *user)
	return nil
}

func (memory *memory) ListAllUsers(list *[]models.User) {
	*list = memory.users
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
