package mysql

import (
	models "CardozoCasariegoLuciano/StudyNotes/Models"
	"CardozoCasariegoLuciano/StudyNotes/configuration"
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var once sync.Once

func getDataBase() *gorm.DB {
	once.Do(func() {
		fmt.Println("Crea la coneccion con la DB")
		db = newDataBase()
	})
	return db
}

func newDataBase() (db *gorm.DB) {
	var err error
	db, err = gorm.Open(mysql.Open(getDataBaseURI()), &gorm.Config{})
	if err != nil {
		fmt.Println("Error en la conexion", err)
		panic(err)
	} else {
		fmt.Println("Conexon de la base de datos exitosa")
		models.UsersMigration(db)
		return db
	}
}

func getDataBaseURI() string {
	config := configuration.GetConfig()

	URI := fmt.Sprintf(
		"root:%s@tcp(%s)/%s?parseTime=true",
		config.DB.Password,
		config.DB.Host,
		config.DB.Name,
	)
	return URI
}
