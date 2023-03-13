package repositories

import (
	"github.com/oxxi/eshop/pkg/entity"
	"gorm.io/gorm"
)

type IProductRepository interface {
	Save(product *entity.ProductoEntity) (*entity.ProductoEntity, error)
	FindAll() ([]entity.ProductoEntity, error)
	FindBy(Id uint64) (*entity.ProductoEntity, error)
	Delete(product *entity.ProductoEntity) error
	Update(Id uint64, product *entity.ProductoEntity) (*entity.ProductoEntity, error)
}

type productRepository struct {
	Db *gorm.DB
}

// Update implements IProductRepository
func (p *productRepository) Update(Id uint64, product *entity.ProductoEntity) (*entity.ProductoEntity, error) {

	updateEntity := entity.ProductoEntity{}

	result := p.Db.Model(updateEntity).Where("id = ?", Id).Updates(product)

	if result.Error != nil {
		return product, result.Error
	}

	return p.FindBy(Id)

}

// Delete implements IProductRepository

func (p *productRepository) Delete(product *entity.ProductoEntity) error {
	err := p.Db.Delete(product).Error
	return err
}

// FindAll implements IProductRepository

func (p *productRepository) FindAll() ([]entity.ProductoEntity, error) {

	var products []entity.ProductoEntity
	var product entity.ProductoEntity
	err := p.Db.Model(&product).Find(&products).Error

	return products, err
}

// FindBy implements IProductRepository

func (p *productRepository) FindBy(Id uint64) (*entity.ProductoEntity, error) {
	var product entity.ProductoEntity
	err := p.Db.First(&product, Id).Error
	return &product, err
}

// Save implements IProductRepository

func (p *productRepository) Save(product *entity.ProductoEntity) (*entity.ProductoEntity, error) {
	result := p.Db.Create(&product)
	return product, result.Error
}

func NewProductRepository(db *gorm.DB) IProductRepository {
	return &productRepository{
		Db: db,
	}
}
