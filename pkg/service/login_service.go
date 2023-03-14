package service

import (
	"github.com/oxxi/eshop/pkg/entity"
	"github.com/oxxi/eshop/pkg/models"
	"github.com/oxxi/eshop/pkg/repositories"
	"golang.org/x/crypto/bcrypt"
)

type ILoginService interface {
	LogIn(user *models.LoginUserModel) bool
	RegisterUSer(user models.RegisterUserModel) error
}

type loginService struct {
	loginRepository repositories.IUserRepository
}

// RegisterUSer implements ILoginService
func (s *loginService) RegisterUSer(user models.RegisterUserModel) error {

	entityUser := entity.ToRegisterUserEnity(user)

	_, err := s.loginRepository.Save(&entityUser)

	return err

}

// LogIn implements ILoginService
func (s *loginService) LogIn(user *models.LoginUserModel) bool {

	//get user
	entityUser, err := s.loginRepository.GetByUser(user.Name)
	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword(entityUser.Password, []byte(user.Password))
	return err == nil
}

var instanceLogin *loginService

func NewLoginService(r repositories.IUserRepository) ILoginService {
	once.Do(func() {
		instanceLogin = &loginService{
			loginRepository: r,
		}
	})

	return instanceLogin
}
