package database

import "context"

func (m *Mysql) UpdateStatuses(ctx context.Context) error {
	stmt, err := m.db.PrepareContext(ctx, "UPDATE shop_payable SET status = DEFERRED"+
		" WHERE status = DEBTOR AND DATE_FORMAT(clear_date, '%Y-%m-%d') < CURDATE() ")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(ctx)
	if err != nil {
		return err
	}

	stm, err := m.db.PrepareContext(ctx, "UPDATE shop_receivable SET status = DEFERRED"+
		" WHERE status = DEBTOR AND DATE_FORMAT(clear_date, '%Y-%m-%d') < CURDATE() ")
	if err != nil {
		return err
	}

	defer stm.Close()

	_, err = stm.Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
