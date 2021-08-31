package database

import (
	"context"
	"github.com/tahmooress/motor-shop/internal/entities/models"
)

func (m *Mysql) UpdateStatuses(ctx context.Context) error {
	stmt, err := m.db.PrepareContext(ctx, "UPDATE shop_payable SET status = ?"+
		" WHERE status = ? AND DATE_FORMAT(clear_date, '%Y-%m-%d') < CURDATE() ")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, models.DEFERRED, models.DEBTOR)
	if err != nil {
		return err
	}

	stm, err := m.db.PrepareContext(ctx, "UPDATE shop_receivable SET status = ?"+
		" WHERE status = ? AND DATE_FORMAT(clear_date, '%Y-%m-%d') < CURDATE() ")
	if err != nil {
		return err
	}

	defer stm.Close()

	_, err = stm.ExecContext(ctx, models.DEFERRED, models.DEBTOR)
	if err != nil {
		return err
	}

	return nil
}
