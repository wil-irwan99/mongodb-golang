package usecase

import (
	"golang-mongodb/model"
	"golang-mongodb/repository"
)

type FindProductByCategoryUseCase interface {
	FindPrdtByCat(category string) ([]model.Product, error)
}

type findProductByCategoryUseCase struct {
	repo repository.ProductRepository
}

func (p *findProductByCategoryUseCase) FindPrdtByCat(category string) ([]model.Product, error) {
	return p.repo.FindByCategory(category)
}

func NewFindProductByCategoryUseCase(repo repository.ProductRepository) FindProductByCategoryUseCase {
	return &findProductByCategoryUseCase{repo: repo}
}
