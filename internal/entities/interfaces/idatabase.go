package interfaces

import (
	"context"
	"github.com/tahmooress/motor-shop/internal/entities/models"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtoadmins"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtocustomers"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtogetshopinventory"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtogetshops"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtoshopequity"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtoshoptrades"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtotransactions"
)

type IDatabase interface {
	GetAdminIDByUserName(ctx context.Context, userName string) (*models.ID, error)
	GetAdminByID(ctx context.Context, adminID models.ID) (*models.Admin, error)
	GetAdmins(ctx context.Context, request *dtoadmins.Request) (*dtoadmins.Response, error)
	GetMotorByPelakNumber(ctx context.Context, pelakNumber string) (*models.Motor, error)
	GetShopsList(ctx context.Context, request *dtogetshops.Request) (*dtogetshops.Response, error)
	GetFactorByNumber(ctx context.Context, factorNumber string, shopID models.ID) (*models.Factor, error)
	GetShopInventory(ctx context.Context, request *dtogetshopinventory.Request) (*dtogetshopinventory.Response, error)
	GetCustomers(ctx context.Context, request *dtocustomers.Request) (*dtocustomers.Response, error)
	GetCustomerByID(ctx context.Context, customerID models.ID) (*models.Customer, error)
	GetShopPayables(ctx context.Context, request *dtoshopequity.Request) (*dtoshopequity.Response, error)
	GetShopReceiveable(ctx context.Context, request *dtoshopequity.Request) (*dtoshopequity.Response, error)
	GetShopBuys(ctx context.Context, request *dtoshoptrades.Request) (*dtoshoptrades.Response, error)
	GetShopSells(ctx context.Context, request *dtoshoptrades.Request) (*dtoshoptrades.Response, error)
	GetTransactionByID(ctx context.Context, transactionID models.ID) (*models.Transaction, error)
	GetShopTransactions(ctx context.Context, request *dtotransactions.Request) (*dtotransactions.Response, error)
	CreateAdmin(ctx context.Context, admin models.Admin) (*models.Admin, error)
	CreateShop(ctx context.Context, shopName string) (*models.ID, error)
	CreateBuyFactor(ctx context.Context, factor models.Factor, shopID models.ID) (*models.Factor, error)
	CreateSellFactor(ctx context.Context, factor models.Factor, shopID models.ID) (*models.Factor, error)
	CreateTransaction(ctx context.Context, transaction models.Transaction) (*models.ID, error)
	UpdateAdmin(ctx context.Context, admin models.Admin) (*models.Admin, error)
	UpdateShopReceivable(ctx context.Context, equityID models.ID) (*models.ShopEquity, error)
	UpdateShopPayable(ctx context.Context, equityID models.ID) (*models.ShopEquity, error)
	UpdateStatuses(ctx context.Context) error
	DeleteAdmin(ctx context.Context, admin models.Admin) (*models.Admin, error)
	Close() error
}
