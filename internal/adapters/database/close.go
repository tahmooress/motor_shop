package database

func (m *Mysql) Close() error {
	return m.db.Close()
}
