package services

import (
	"mvc/domain"
	"mvc/utils"
)

func GetUser(userId int64)  (*domain.User,*utils.ApplicationError) {
		return domain.UserDao.GetUser(userId)
	
}