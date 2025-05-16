package repository

import (
	"database/sql"

	"github.com/nyakovchuk/vsch_youtube_bot/internal/domain/platform"
	"github.com/nyakovchuk/vsch_youtube_bot/internal/domain/user"
)

type Repository struct {
	UserRepository     user.Repository
	PlatformRepository platform.Repository
}

func New(db *sql.DB) *Repository {
	return &Repository{
		UserRepository:     user.NewRepository(db),
		PlatformRepository: platform.NewRepository(db),
	}
}
