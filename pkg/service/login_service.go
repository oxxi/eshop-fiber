package service

import "github.com/oxxi/eshop/pkg/models"

type ILoginService interface {
	LogIn(user *models.LoginUserModel) bool
	RegisterUSer(user *models.RegisterUserModel)
}

type loginService struct {
}

// RegisterUSer implements ILoginService
func (s *loginService) RegisterUSer(user *models.RegisterUserModel) {
	panic("unimplemented")
}

// LogIn implements ILoginService
func (s *loginService) LogIn(user *models.LoginUserModel) bool {
	return true
}

var instanceLogin *loginService

func NewLoginService() ILoginService {
	once.Do(func() {
		instanceLogin = &loginService{}
	})

	return instanceLogin
}
