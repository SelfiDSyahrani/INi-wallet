package manager

import "INi-Wallet/repository"

type RepositoryManger interface {
	UserRepository() repository.UserRepository
	// transactionRepository() repository.transactionRepository
}

type repositoryManger struct {
	infra InfraManager
}

func (r *repositoryManger) UserRepository() repository.UserRepository {
	return repository.NewUserRepository(r.infra.SqlDb())
}

func NewRepositoryManager(infraManager InfraManager) RepositoryManger {
	return &repositoryManger{
		infra: infraManager,
	}
}
