package repository

import (
	"nexample/internal/database"

	"gorm.io/gorm"
)

type Repo[T any] struct {
	db *gorm.DB
}

func New[T any]() *Repo[T] {
	return &Repo[T]{db: database.DB}
}

func (r *Repo[T]) Create(entity *T) error {
	return r.db.Create(entity).Error
}

func (r *Repo[T]) GetByID(id uint) (*T, error) {
	var entity T
	err := r.db.First(&entity, id).Error
	return &entity, err
}

func (r *Repo[T]) GetAll() ([]T, error) {
	var entities []T
	err := r.db.Find(&entities).Error
	return entities, err
}

func (r *Repo[T]) Update(entity *T) error {
	return r.db.Save(entity).Error
}

func (r *Repo[T]) Delete(id uint) error {
	var entity T
	return r.db.Delete(&entity, id).Error
}

func (r *Repo[T]) Where(query string, args ...any) ([]T, error) {
	var entities []T
	err := r.db.Where(query, args...).Find(&entities).Error
	return entities, err
}

func (r *Repo[T]) FirstWhere(query string, args ...any) (*T, error) {
	var entity T
	err := r.db.Where(query, args...).First(&entity).Error
	return &entity, err
}

func (r *Repo[T]) Preload(preload string) *gorm.DB {
	return r.db.Preload(preload)
}

func (r *Repo[T]) DB() *gorm.DB {
	return r.db
}
