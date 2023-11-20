package repository

import models "CardozoCasariegoLuciano/StudyNotes/Models"

type IStorage interface {
	SaveUser(user *models.User) error
	ListAllUsers(*[]models.User) error
	FindUserByEmail(email string) models.User
	GetUserByID(id uint) models.User
}
