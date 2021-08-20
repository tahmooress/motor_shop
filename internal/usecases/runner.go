package usecases

import (
	"context"
	"time"

	"github.com/tahmooress/motor-shop/internal/pkg/logger"
)

func (u *UseCases) runner(ctx context.Context, logger *logger.Logger) {
	ticker := time.NewTicker(time.Hour * time.Duration(24))

	for range ticker.C {
		err := u.IDatabase.UpdateStatuses(ctx)
		if err != nil {
			logger.Error.Print(err)
		}
	}
}
