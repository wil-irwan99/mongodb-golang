package usecase

import (
	"golang-mongodb/model"
	"golang-mongodb/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FindProductByIdUseCase interface {
	FindPrdtById(id primitive.ObjectID) (model.Product, error)
}

type findProductByIdUseCase struct {
	repo repository.ProductRepository
}

func (p *findProductByIdUseCase) FindPrdtById(id primitive.ObjectID) (model.Product, error) {
	return p.repo.FindById(id)
}

func NewFindProductByIdUseCase(repo repository.ProductRepository) FindProductByIdUseCase {
	return &findProductByIdUseCase{repo: repo}
}
