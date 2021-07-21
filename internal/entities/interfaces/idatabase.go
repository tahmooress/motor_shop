package interfaces

import (
	"context"
	"github.com/tahmooress/motor-shop/internal/entities/models"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtoadmins"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtobuy"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtodeleteadmin"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtogetshopinventory"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtogetshops"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtosell"
)

type IDatabase interface {
	GetAdminIDByUserName(ctx context.Context, userName string) (*models.ID, error)
	GetAdminByID(ctx context.Context, adminID models.ID) (*models.Admin, error)
	GetAdminAccessibility(ctx context.Context, adminID models.ID) ([]string, error)
	GetAdmins(ctx context.Context, request *dtoadmins.Request) (*dtoadmins.Response, error)
	GetShopsList(ctx context.Context, request *dtogetshops.Request) (*dtogetshops.Response, error)
	GetBuyFactorByNumber(ctx context.Context, factorNumber string) (*dtobuy.Response, error)
	GetShopInventory(ctx context.Context, request *dtogetshopinventory.Request) (*dtogetshopinventory.Response, error)
	GetSellFactorByNumber(ctx context.Context, factorNumber string) (*dtosell.Response, error)
	CreateAdmin(ctx context.Context, userName, hashedPassword string, accessibility []models.ID) (*models.Admin, error)
	CreateBuyFactor(ctx context.Context, request *dtobuy.Request) (string, error)
	CreateSellFactor(ctx context.Context, request *dtosell.Request) (string, error)
	UpdateAdmin(ctx context.Context, admin models.Admin) (*models.Admin, error)
	DeleteAdmin(ctx context.Context, request *dtodeleteadmin.Request) (*dtodeleteadmin.Response, error)
	Close() error
}
