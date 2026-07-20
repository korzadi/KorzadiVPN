package database

func ConfigureSQLite() {

	DB.Exec(`
	PRAGMA journal_mode=WAL;
	`)

	DB.Exec(`
	PRAGMA busy_timeout=5000;
	`)
}
