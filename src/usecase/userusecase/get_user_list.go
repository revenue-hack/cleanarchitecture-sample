package userusecase

import "github.com/revenue-hack/cleanarchitecture-sample/src/entity"

type GetUserListUsecase struct {
	userRepository UserRepository
}

func NewGetUserList(userRepo UserRepository) *GetUserListUsecase {
	return &GetUserListUsecase{
		userRepository: userRepo,
	}
}

func (use *GetUserListUsecase) Exec() ([]*entity.User, error) {
	return use.userRepository.FindAll()
}
