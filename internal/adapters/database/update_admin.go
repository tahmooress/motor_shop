package database

import (
	"context"
	"fmt"
	"strings"

	"github.com/tahmooress/motor-shop/internal/entities/models"
)

func (m *Mysql) UpdateAdmin(ctx context.Context, admin models.Admin) (*models.Admin, error) {
	tx, err := m.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("mysql >> UpdateAdmin() >> db.Begin() >> %w", err)
	}
	stmt, err := tx.PrepareContext(ctx, "UPDATE admin_users SET user_name = ?, password = ? WHERE id = ?")
	if err != nil {
		return nil, fmt.Errorf("mysql >> UpdateAdmin() >> PrepareContext() >> %w", err)
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, admin.UserName, admin.Password, admin.ID)
	if err != nil {
		return nil, fmt.Errorf("mysql >> UpdateAdmin() >> ExecContext() >> %w", err)
	}

	deleteStmt, err := tx.PrepareContext(ctx, "DELETE FROM accessibility WHERE admin_id = ?")
	if err != nil {
		return nil, fmt.Errorf("mysql >> UpdateAdmin() >> PrepareContext() >> %w", err)
	}

	_, err = deleteStmt.ExecContext(ctx, admin.ID)
	if err != nil {
		return nil, fmt.Errorf("mysql >> UpdateAdmin >> ExecContext() >> %w", err)
	}

	if admin.Accessibility == nil {
		err = tx.Commit()
		if err != nil {
			return nil, fmt.Errorf("mysql >> UpdateAdmin() >> tx.Commit() >> %w", err)
		}

		return &models.Admin{
			ID:            admin.ID,
			UserName:      admin.UserName,
			Password:      admin.Password,
			Accessibility: nil,
		}, nil
	}

	rawStmt := "INSERT INTO accessibility(id,admin_id, shop_id) VALUES "
	valStmt := ""
	args := make([]interface{}, 0)

	for _, id := range admin.Accessibility {
		valStmt += "(?,?,?),"
		args = append(args, models.NewID(), admin.ID, id)
	}

	rawStmt += strings.TrimRight(valStmt, ",")

	insertStmt, err := tx.PrepareContext(ctx, rawStmt)
	if err != nil {
		tx.Rollback()

		return nil, fmt.Errorf("mysql >> UpdateAdmin() >> PrepareContext() >> %w", err)
	}

	_, err = insertStmt.ExecContext(ctx, args...)
	if err != nil {
		tx.Rollback()

		return nil, fmt.Errorf("mysql >> UpdateAdmin() >> ExecContext() >> %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("mysql >> UpdateAdmin() >> tx.Commit() >> %w", err)
	}

	result, err := m.GetAdminByID(ctx, admin.ID)
	if err != nil {
		return nil, fmt.Errorf("mysql >> UpdateAdmin() >> %w", err)
	}

	return result, nil
}
