package database

import (
	"context"
	"fmt"
	"github.com/tahmooress/motor-shop/internal/entities/models"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtogetshopinventory"
	"math"
)

func (m *Mysql) GetShopInventory(ctx context.Context,
	request *dtogetshopinventory.Request) (*dtogetshopinventory.Response, error) {
	stmt, err := m.db.PrepareContext(ctx, "SELECT shop_inventory.id, factor_number, motors.id,"+
		" motors.model_name, motors.pelak_number, motors.body_number,"+
		" motors.color, motors.model_year,"+
		" motors.created_at, motors.updated_at,"+
		" shop_inventory.created_at, shop_inventory.updated_at"+
		" FROM shop_inventory INNER JOIN motors ON shop_inventory.motor_id = motors.id "+
		"WHERE shop_inventory.shop_id = ? LIMIT ? OFFSET ?")
	if err != nil {
		return nil, fmt.Errorf("mysql >> GetShopInventory >> PrepareContext() >> %w", err)
	}

	defer stmt.Close()

	var response dtogetshopinventory.Response

	if request.Query.Size == 0 {
		request.Query.Size = 10
	}

	rows, err := stmt.QueryContext(ctx, request.ShopID, request.Query.Size, (request.Number-1)*request.Size)
	if err != nil {
		return nil, fmt.Errorf("mysql >> GetShopInventory >> QueryContext() >> %w", err)
	}

	defer rows.Close()

	response.Data = make([]models.Inventory, 0)

	for rows.Next() {
		var temp models.Inventory

		err = rows.Scan(&temp.ID, &temp.FactorNumber, &temp.Motor.ID,
			&temp.Motor.ModelName, &temp.Motor.PelakNumber, &temp.Motor.BodyNumber,
			&temp.Motor.Color, &temp.Motor.ModelYear, &temp.Motor.CreatedAt,
			&temp.Motor.UpdatedAt, &temp.CreatedAt, &temp.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("mysql >> GetShopInventory >> rows.Scan() >> %w", err)
		}

		response.Data = append(response.Data, temp)
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("mysql >> GetShopInventory >> rows.Err() >> %w", err)
	}

	metaStmt, err := m.db.PrepareContext(ctx, "SELECT count(*) FROM shop_inventory WHERE shop_id = ?")
	if err != nil {
		return nil, fmt.Errorf("mysql >> GetShopInventory >> PrepareContext() >> %w", err)
	}

	defer metaStmt.Close()

	err = metaStmt.QueryRowContext(ctx, request.ShopID).Scan(&response.Meta.Total)
	if err != nil {
		return nil, fmt.Errorf("mysql >> GetShopInventory >> QueryRowContext() >> %w", err)
	}

	response.Meta.PageSize = request.Size

	response.Meta.LastPage = int(math.Ceil(float64(response.Meta.Total) / float64(response.Meta.PageSize)))
	if response.Meta.LastPage == 0 {
		response.Meta.LastPage = 1
	}

	response.Meta.PageSize = request.Size
	response.Meta.PageNumber = request.Number

	return &response, nil
}
