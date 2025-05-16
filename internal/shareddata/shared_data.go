package shareddata

import (
	"context"
	"fmt"

	"github.com/nyakovchuk/vsch_youtube_bot/config"
	"github.com/nyakovchuk/vsch_youtube_bot/internal/domain/platform"
	"github.com/nyakovchuk/vsch_youtube_bot/internal/service"
)

type Data struct {
	Platform platform.Platform
}

func New(ctx context.Context, cfg *config.Config, services *service.Service) Data {
	platform, err := services.Platform.GetByName(ctx, cfg.Platform)
	if err != nil {
		fmt.Println("error getting platform", err)
	}

	return Data{
		Platform: platform,
	}
}
