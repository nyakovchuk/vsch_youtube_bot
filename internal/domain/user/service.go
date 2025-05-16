package user

import (
	"context"

	"github.com/nyakovchuk/vsch_youtube_bot/internal/apperrors"
)

type Service interface {
	Register(ctx context.Context, platformId int, user User) error
	IsRegistered(platformId int, externalId string) (bool, error)
}

type service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &service{
		repo: repository,
	}
}

func (s *service) Register(ctx context.Context, platformId int, user User) error {

	repoUserDto := ToDto(user)
	repoUserDto.PlatformID = platformId

	if err := s.repo.RegisterUser(ctx, repoUserDto); err != nil {
		return apperrors.Wrap(apperrors.ErrUserRegistration, err)
	}

	return nil
}

func (s *service) IsRegistered(platformId int, externalId string) (bool, error) {
	exist, err := s.repo.IsRegistered(platformId, externalId)
	if err != nil {
		return false, err
	}

	if !exist {
		return false, nil
	}

	return true, nil
}
