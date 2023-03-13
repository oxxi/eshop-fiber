package repositories

import (
	"github.com/oxxi/eshop/pkg/entity"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Save(u *entity.UserEntity) (entity.UserEntity, error)
	GetAll() ([]entity.UserEntity, error)
	GetByUser(userName string) (entity.UserEntity, error)
	GetById(Id uint64) (entity.UserEntity, error)
	Update(Id uint64, user *entity.UserEntity) (entity.UserEntity, error)
	Delete(user *entity.UserEntity) error
}

type userRepository struct {
	Db *gorm.DB
}

// Delete implements IUserRepository
func (r *userRepository) Delete(user *entity.UserEntity) error {
	panic("unimplemented")
}

// GetAll implements IUserRepository
func (r *userRepository) GetAll() ([]entity.UserEntity, error) {
	panic("unimplemented")
}

// GetById implements IUserRepository
func (r *userRepository) GetById(Id uint64) (entity.UserEntity, error) {
	panic("unimplemented")
}

// GetByUser implements IUserRepository
func (r *userRepository) GetByUser(userName string) (entity.UserEntity, error) {
	panic("unimplemented")
}

// Save implements IUserRepository
func (r *userRepository) Save(u *entity.UserEntity) (entity.UserEntity, error) {
	panic("unimplemented")
}

// Update implements IUserRepository
func (r *userRepository) Update(Id uint64, user *entity.UserEntity) (entity.UserEntity, error) {
	panic("unimplemented")
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{
		Db: db,
	}
}
