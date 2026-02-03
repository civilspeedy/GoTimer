package sqlTime

const (
	Create = `
	CREATE TABLE IF NOT EXISTS time(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		date TEXT UNIQUE,
		time INTEGER
	);`
	Drop        = "DROP TABLE IF EXISTS time;"
	SelectWhere = "SELECT time FROM time WHERE date = ?;"
	SelectAll   = "SELECT * FROM time;"
	Insert      = "INSERT INTO time(date, time) VALUES(?, ?);"
	Update      = "UPDATE time SET time = ? WHERE date = ?;"
)
