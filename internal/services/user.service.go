package services

import "github.com/longtk26/go-ecommerce/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

func (us *UserService) GetUserByID() string {
	return us.userRepo.GetUserByID()
}