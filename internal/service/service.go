package service

import (
	"github.com/nyakovchuk/vsch_youtube_bot/internal/domain/platform"
	"github.com/nyakovchuk/vsch_youtube_bot/internal/domain/user"
	"github.com/nyakovchuk/vsch_youtube_bot/internal/repository"
)

type Service struct {
	User     user.Service
	Platform platform.Service
}

func New(repo *repository.Repository) *Service {
	userService := user.NewService(repo.UserRepository)
	platformService := platform.NewService(repo.PlatformRepository)

	return &Service{
		User:     userService,
		Platform: platformService,
	}
}
