package usecase

import (
	"golang-mongodb/model"
	"golang-mongodb/repository"
)

type FindAllProductUseCase interface {
	FindAll(page int64, limit int64) ([]model.Product, error)
}

type findAllProductUseCase struct {
	repo repository.ProductRepository
}

func (p *findAllProductUseCase) FindAll(page int64, limit int64) ([]model.Product, error) {
	return p.repo.Retrieve(page, limit)
}

func NewFindAllProductUseCase(repo repository.ProductRepository) FindAllProductUseCase {
	return &findAllProductUseCase{repo: repo}
}
