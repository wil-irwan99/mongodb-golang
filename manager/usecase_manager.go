package manager

import "golang-mongodb/usecase"

type UseCaseManager interface {
	ProductRegistrationUseCase() usecase.ProductRegistrationUseCase
	FindAllProductUseCase() usecase.FindAllProductUseCase
	FindProductById() usecase.FindProductByIdUseCase
	FindProductByCategory() usecase.FindProductByCategoryUseCase
	UpdateProductById() usecase.ProductUpdateUseCase
	DeleteProductById() usecase.ProductDeleteUseCase
}

type useCaseManager struct {
	repoManager RepositoryManager
}

func (u *useCaseManager) ProductRegistrationUseCase() usecase.ProductRegistrationUseCase {
	return usecase.NewProductRegistrationUseCase(u.repoManager.ProductRepo())
}

func (u *useCaseManager) FindAllProductUseCase() usecase.FindAllProductUseCase {
	return usecase.NewFindAllProductUseCase(u.repoManager.ProductRepo())
}

func (u *useCaseManager) FindProductById() usecase.FindProductByIdUseCase {
	return usecase.NewFindProductByIdUseCase(u.repoManager.ProductRepo())
}

func (u *useCaseManager) FindProductByCategory() usecase.FindProductByCategoryUseCase {
	return usecase.NewFindProductByCategoryUseCase(u.repoManager.ProductRepo())
}

func (u *useCaseManager) UpdateProductById() usecase.ProductUpdateUseCase {
	return usecase.NewProductUpdateUseCase(u.repoManager.ProductRepo())
}

func (u *useCaseManager) DeleteProductById() usecase.ProductDeleteUseCase {
	return usecase.NewProductDeleteUseCase(u.repoManager.ProductRepo())
}

func NewUseCaseManager(repoManager RepositoryManager) UseCaseManager {
	return &useCaseManager{repoManager: repoManager}
}
