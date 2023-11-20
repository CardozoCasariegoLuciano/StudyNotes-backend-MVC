package mysql

import (
	models "CardozoCasariegoLuciano/StudyNotes/Models"
	repository "CardozoCasariegoLuciano/StudyNotes/Repository"
	"sync"

	"gorm.io/gorm"
)

var db_MySQL *database
var once_mySQL sync.Once

type database struct {
	db *gorm.DB
}

func NewDataBase() repository.IStorage {
	once_mySQL.Do(func() {
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

func (st *database) GetUserByID(id uint) models.User {
	user := models.User{}
	st.db.Where("id = ?", id).First(&user)
	return user
}

func (st *database) EditUser(id uint, name string, image string) {
	user := models.User{}
	st.db.Model(user).Where("id = ?", id).Updates(
		map[string]interface{}{"name": name, "image": image},
	)
}
