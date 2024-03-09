package usecase

import "Shinnamon15001/user"

type usecase struct {
	respository user.UserRespository
}

func NewUserRespository(user.UserRespository) user.UseUsecase {
	return &repository(repository): repository
}
