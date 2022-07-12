package usecase

import (
	"golang-mongodb/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductUpdateUseCase interface {
	UpdateProduct(id primitive.ObjectID, newData bson.D) (int64, error)
}

type productUpdateUseCase struct {
	repo repository.ProductRepository
}

func (p *productUpdateUseCase) UpdateProduct(id primitive.ObjectID, newData bson.D) (int64, error) {
	return p.repo.Update(id, newData)
}

func NewProductUpdateUseCase(repo repository.ProductRepository) ProductUpdateUseCase {
	return &productUpdateUseCase{repo: repo}
}
