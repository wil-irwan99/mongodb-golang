package usecase

import (
	"golang-mongodb/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductDeleteUseCase interface {
	DeleteProduct(id primitive.ObjectID) (int64, error)
}

type productDeleteUseCase struct {
	repo repository.ProductRepository
}

func (p *productDeleteUseCase) DeleteProduct(id primitive.ObjectID) (int64, error) {
	return p.repo.Delete(id)
}

func NewProductDeleteUseCase(repo repository.ProductRepository) ProductDeleteUseCase {
	return &productDeleteUseCase{repo: repo}
}
