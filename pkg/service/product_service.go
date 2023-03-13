package service

import (
	"sync"

	"github.com/oxxi/eshop/pkg/entity"
	"github.com/oxxi/eshop/pkg/models"
	"github.com/oxxi/eshop/pkg/repositories"
)

type IProductService interface {
	Save(m models.ProductModel) (models.ProductModel, error)
	GetAll() ([]models.ProductModel, error)
	GetBy(Id uint64) (models.ProductModel, error)
	Delete(Id uint64) error
	Update(Id uint64, model models.ProductModel) (models.ProductModel, error)
}

type productService struct {
	repository repositories.IProductRepository
}

// Update implements IProductService
func (s *productService) Update(Id uint64, model models.ProductModel) (models.ProductModel, error) {

	updateEntity := entity.ToEntity(model)
	updateEntity.Id = uint(Id)
	updatedEntity, err := s.repository.Update(Id, &updateEntity)
	if err != nil {
		return models.ProductModel{}, err
	}

	return updatedEntity.FromEntity(), nil
}

// Delete implements IProductService
func (s *productService) Delete(Id uint64) error {
	productEntity, err := s.repository.FindBy(Id)
	if err != nil {
		return err
	}

	s.repository.Delete(productEntity)

	return nil
}

// GetAll implements IProductService
func (s productService) GetAll() ([]models.ProductModel, error) {
	productEntity, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	var products []models.ProductModel

	for _, item := range productEntity {
		products = append(products, item.FromEntity())
	}
	return products, nil
}

// GetBy implements IProductService
func (s *productService) GetBy(Id uint64) (models.ProductModel, error) {
	var product models.ProductModel
	entity, err := s.repository.FindBy(Id)
	if err != nil {
		return product, err
	}
	product = entity.FromEntity()
	return product, nil
}

// Save implements IProductService
func (s *productService) Save(m models.ProductModel) (models.ProductModel, error) {
	productEntity := entity.ToEntity(m)
	e, err := s.repository.Save(&productEntity)
	return e.FromEntity(), err
}

var once sync.Once
var instance *productService

func NewProductService(r repositories.IProductRepository) IProductService {
	once.Do(func() {
		instance = &productService{
			repository: r,
		}
	})
	return instance
}
