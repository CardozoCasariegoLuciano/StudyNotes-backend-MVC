package repository

import models "CardozoCasariegoLuciano/StudyNotes/Models"

type IStorage interface {
	Save(user models.User) models.User
	ListAll() []models.User
}
