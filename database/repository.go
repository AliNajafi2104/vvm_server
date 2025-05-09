package database

import "gorm.io/gorm"

type Database[T any] interface {
	FindByID(id interface{}) (*T, error)
	CreateEntity(entity *T) error
	UpdateEntity(entity *T) error
	DeleteByID(id interface{}) error
	FindAll() ([]T, error)
}

type GormDatabase[T any] struct {
	DB *gorm.DB
}

func (g *GormDatabase[T]) FindByID(id interface{}) (*T, error) {
	var entity T
	result := g.DB.First(&entity, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &entity, nil
}

func (g *GormDatabase[T]) CreateEntity(entity *T) error {

	result := g.DB.Create(entity)
	return result.Error

}

func (g *GormDatabase[T]) UpdateEntity(entity *T) error {

	result := g.DB.Save(entity)
	return result.Error

}

func (g *GormDatabase[T]) DeleteByID(id interface{}) error {
	var entity T
	result := g.DB.Delete(&entity, id)
	return result.Error
}

func (g *GormDatabase[T]) FindAll() ([]T, error) {
	var entities []T
	result := g.DB.Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return entities, nil
}
