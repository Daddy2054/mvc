package services

import (
	"mvc/domain"
	"mvc/utils"
)

type usersService struct{}

var (
	UsersService usersService
)

func (u *usersService) GetUser(userId int64)  (*domain.User,*utils.ApplicationError) {
		return domain.UserDao.GetUser(userId)
	
}