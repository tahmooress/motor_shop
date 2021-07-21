package usecases

import (
	"context"

	"github.com/tahmooress/motor-shop/internal/entities/interfaces"
	"github.com/tahmooress/motor-shop/internal/pkg/logger"
)

const SECRETKEY = "secret"

type UseCases struct {
	IDatabase interfaces.IDatabase
}

func New(ctx context.Context, IDatabase interfaces.IDatabase, logger *logger.Logger) (interfaces.IUseCases, error) {
	return &UseCases{
		IDatabase: IDatabase,
	}, nil
}
