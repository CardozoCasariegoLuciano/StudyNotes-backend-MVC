package repository

import models "CardozoCasariegoLuciano/StudyNotes/Models"

type IMemory interface {
	Save(user models.User) models.User
	ListAll() []models.User
}
