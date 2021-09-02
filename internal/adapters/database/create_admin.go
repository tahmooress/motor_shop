package database

import (
	"context"
	"fmt"
	"strings"

	"github.com/tahmooress/motor-shop/internal/entities/models"
)

func (m *Mysql) CreateAdmin(ctx context.Context, admin models.Admin) (*models.Admin, error) {
	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("mysql >> CreateAdmin() >> db.Begin() >> %w", err)
	}

	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, "INSERT IGNORE INTO admin_users(id, user_name, password)"+
		"VALUES(?,?,?)")
	if err != nil {
		return nil, fmt.Errorf("mysql >> CreateAdmin() >> PrepareContext() >> %w", err)
	}

	defer stmt.Close()

	adminID := models.NewID()

	result, err := stmt.ExecContext(ctx, adminID, admin.UserName, admin.Password)
	if err != nil {
		return nil, fmt.Errorf("mysql >> CreateAdmin() >> ExecContext() >> %w", err)
	}

	n, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("mysql >> CreateAdmin >> RowsAffected() >> %w", err)
	}

	if n == 0 {
		return nil, models.ErrUserIsTaken
	}

	statement := "INSERT INTO accessibility(id, admin_id, shop_id) VALUES"
	valStmts := ""
	args := make([]interface{}, 0)

	for _, shop := range admin.Shops {
		valStmts += "(?,?,?),"
		args = append(args, models.NewID(), adminID, shop.ID)
	}

	statement += strings.TrimRight(valStmts, ",")

	insertStmt, err := tx.PrepareContext(ctx, statement)
	if err != nil {
		_ = tx.Rollback()

		return nil, fmt.Errorf("mysql >> CreateAdmin() PreapareContext() >> %w", err)
	}

	_, err = insertStmt.ExecContext(ctx, args...)
	if err != nil {
		_ = tx.Rollback()

		return nil, fmt.Errorf("mysql >> CreateAdmin() >> ExecContext() >> %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("mysql >> CreateAdmin() >> tx.Commit() >> %w", err)
	}

	respAdmin, err := m.GetAdminByID(ctx, adminID)
	if err != nil {
		return nil, fmt.Errorf("mysql >> CreateAdmin() >> %w", err)
	}

	return respAdmin, nil
}
