package service

import "github.com/xuanvinh9411/go-ecommerce-backend/internal/repo"

type UserService struct { 
	UserRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) GetInfoUser() string {
	return "name: xuanvinh9411, old: 18"
}