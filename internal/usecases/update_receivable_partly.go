package usecases

import (
	"context"
	"errors"

	"github.com/tahmooress/motor-shop/internal/entities/models"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtoupdateequities"
)

func (u *UseCases) UpdateReceivablePartly(ctx context.Context,
	request *dtoupdateequities.Request) (*models.ShopEquity, error) {
	shopEquity, shopID, err := u.IDatabase.GetReceivableByID(ctx, request.ID)
	if err != nil {
		return nil, err
	}

	if shopEquity.Amount == request.PayedAmount {
		return u.IDatabase.UpdateShopPayable(ctx, request.ID)
	}

	if shopEquity.Amount < request.PayedAmount {
		return nil, errors.New("cant update for payed amount bigger thant equity amount")
	}

	newEquityID, err := u.IDatabase.UpdateReceivablePartly(ctx, *shopEquity, *shopID, request)
	if err != nil {
		return nil, err
	}

	newEquity, _, err := u.IDatabase.GetReceivableByID(ctx, *newEquityID)
	if err != nil {
		return nil, err
	}

	return newEquity, nil
}




