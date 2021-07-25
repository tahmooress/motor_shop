package database

import (
	"context"
	"fmt"
	"math"

	"github.com/tahmooress/motor-shop/internal/entities/models"
	"github.com/tahmooress/motor-shop/internal/pkg/query"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtoadmins"
)

func (m *Mysql) GetAdmins(ctx context.Context, request *dtoadmins.Request) (*dtoadmins.Response, error) {
	response := dtoadmins.Response{
		Admins: make([]models.Admin, 0),
		Meta: query.Meta{
			PageNumber: request.Number,
			PageSize:   request.Size,
		},
	}

	countStmt, err := m.db.PrepareContext(ctx, "SELECT count(DISTINCT id) FROM admin_users")
	if err != nil {
		return nil, fmt.Errorf("mysql >> GetAdmins >> PrepareContext() >> %w", err)
	}

	defer countStmt.Close()

	err = countStmt.QueryRowContext(ctx).Scan(&response.Meta.Total)
	if err != nil {
		return nil, fmt.Errorf("mysql >> GetAdmins >> QueryRowContext() >> %w", err)
	}

	stmt, err := m.db.PrepareContext(ctx, "SELECT id,user_name,password,created_at,updated_at FROM admin_users LIMIT ? OFFSET ?")
	if err != nil {
		return nil, fmt.Errorf("mysql >> GetAdmins >> PrepareContext() >> %w", err)
	}

	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, request.Size, (request.Number-1)*request.Size)
	if err != nil {
		return nil, fmt.Errorf("mysql >> GetAdmins >> QueryContext() >> %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var admin models.Admin

		err = rows.Scan(&admin.ID, &admin.UserName, &admin.Password, &admin.CreatedAt, &admin.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("mysql >> GetAdmin >> Scan() >> %w", err)
		}

		response.Admins = append(response.Admins, admin)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("mysql >> GetAdmin >> rwos.Err() >> %w", err)
	}

	stm, err := m.db.PrepareContext(ctx, "SELECT shops.id, shops.shop_name, shops.created_at, shops.updated_at FROM "+
		"shops LEFT JOIN accessibility ON shops.id = accessibility.shop_id WHERE accessibility.admin_id = ?")
	if err != nil {
		return nil, fmt.Errorf("mysql >> PrepareContext() >> %w", err)
	}

	defer stm.Close()

	for index, admin := range response.Admins {
		rows, err := stm.QueryContext(ctx, admin.ID)
		if err != nil {
			rows.Close()

			return nil, fmt.Errorf("mysql >> GetAdmins >>  QueryContext() >> %w", err)
		}

		response.Admins[index].Shops = make([]models.Shop, 0)

		for rows.Next() {
			var shop models.Shop

			err = rows.Scan(&shop.ID, &shop.ShopName, &shop.CreatedAt, &shop.UpdatedAt)
			if err != nil {
				rows.Close()

				return nil, fmt.Errorf("mysql >> GetAdmins >> rows.Scan() >> %w", err)
			}

			response.Admins[index].Shops = append(response.Admins[index].Shops, shop)
		}

		rows.Close()
	}

	response.Meta.LastPage = int(math.Ceil(float64(response.Meta.Total) / float64(response.Meta.PageSize)))
	if response.Meta.LastPage == 0 {
		response.Meta.LastPage = 1
	}

	return &response, nil
}
