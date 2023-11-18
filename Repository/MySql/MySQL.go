package mysql

import (
	models "CardozoCasariegoLuciano/StudyNotes/Models"
	"fmt"
	"sync"

	"gorm.io/gorm"
)

var db_MySQL *database
var once_mySQL sync.Once

type database struct {
	db *gorm.DB
}

func NewDataBase() *database {
	once_mySQL.Do(func() {
		fmt.Println("Pasa por aca MySQL")
		db_MySQL = &database{db: getDataBase()}
	})
	return db_MySQL
}

func (st *database) SaveUser(user *models.User) error {
	gormResp := st.db.Save(&user)
	if gormResp.RowsAffected == 0 {
		return gormResp.Error
	}

	return nil
}

func (st *database) ListAllUsers(list *[]models.User) error {
	gormResp := st.db.Find(list)
	return gormResp.Error
}

func (st *database) FindUserByEmail(email string) models.User {
	user := models.User{}
	st.db.Where("email = ?", email).First(&user)
	return user
}
