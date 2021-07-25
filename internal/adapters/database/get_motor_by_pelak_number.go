package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/tahmooress/motor-shop/internal/entities/models"
)

func (m *Mysql) GetMotorByPelakNumber(ctx context.Context, pelakNumber string) (*models.Motor, error) {
	stmt, err := m.db.PrepareContext(ctx, "SELECT id, model_name, pelak_number, body_number,"+
		" color, model_year, created_at, updated_at FROM motors WHERE pelak_number = ?")
	if err != nil {
		return nil, fmt.Errorf("mysql >> GetMotorByPelakNumber >> db.PrepareContext() >> %w", err)
	}

	defer stmt.Close()

	var motor models.Motor

	err = stmt.QueryRowContext(ctx, pelakNumber).Scan(&motor.ID, &motor.ModelName, &motor.PelakNumber,
		&motor.BodyNumber, &motor.Color, &motor.ModelYear,
		&motor.CreatedAt, &motor.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrMotorIsNotExist
		}
		return nil, fmt.Errorf("mysql >> GetMotorByPelakNumber >> QueryRowContext() >> %w", err)
	}

	return &motor, nil
}
