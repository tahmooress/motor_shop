package interfaces

import (
	"context"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtoadmins"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtobuy"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtodeleteadmin"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtogetshopinventory"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtogetshops"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtosell"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtoupdateadmin"

	"github.com/tahmooress/motor-shop/internal/port/dto/dtoadmin"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtologin"
)

type IUseCases interface {
	//UpdateReceives(ctx context.Context, request dtoupdatereceives.Request) (dtoupdatereceives.Response, error)
	//UpDatePays(ctx context.Context, request dtoupdatepayables.Request) (dtoupdatepayables.Response, error)
	//CreateCustomer(ctx context.Context, request dtocustomer.Request) (dtocustomer.Response, error)
	//GetStocks(ctx context.Context, request dtogetshopinventory.Request) (dtogetshopinventory.Response, error)
	//GetBalances(ctx context.Context, request dtogetbalance.Request) (dtogetbalance.Response, error)
	//GetReceivables(ctx context.Context, request dtoshoprecieveables.Request) (dtoshoprecieveables.Response, error)
	//GetPayables(ctx context.Context, request dtoshoppayables.Request) (dtoshoppayables.Response, error)
	//GetSells(ctx context.Context, request dtogetsell.Request) (dtogetsell.Response, error)
	//GetBuys(ctx context.Context, request dtogetbuy.Request) (dtogetbuy.Response, error)
	//Sell(ctx context.Context, request dtosell.Request) (dtosell.Response, error)
	//Buy(ctx context.Context, request dtobuy.Request) (dtobuy.Response, error)
	//Transaction(ctx context.Context, request dtotransaction.Request) (dtotransaction.Response, error)
	Authentication(ctx context.Context, tokenString string) (context.Context, error)
	CreateAdmin(ctx context.Context, request *dtoadmin.Request) (*dtoadmin.Response, error)
	UpdateAdmin(ctx context.Context, request *dtoupdateadmin.Request) (*dtoupdateadmin.Response, error)
	DeleteAdmin(ctx context.Context, request *dtodeleteadmin.Request) (*dtodeleteadmin.Response, error)
	GetAdmins(ctx context.Context, request *dtoadmins.Request) (*dtoadmins.Response, error)
	GetShopsList(ctx context.Context, request *dtogetshops.Request) (*dtogetshops.Response, error)
	Login(ctx context.Context, request *dtologin.Request) (*dtologin.Response, error)
	Buy(ctx context.Context, request *dtobuy.Request) (*dtobuy.Response, error)
	Sell(ctx context.Context, request *dtosell.Request) (*dtosell.Response, error)
	//GetSells(ctx context.Context, request dtogetsell.Request) (dtogetsell.Response, error)
	GetShopInventory(ctx context.Context,
		request *dtogetshopinventory.Request) (*dtogetshopinventory.Response, error)
}
