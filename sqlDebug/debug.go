package sqlDebug

const (
	Create = `CREATE TABLE IF NOT EXISTS debug(
		id INTEGER PRIMARY KEY AUTOINCREMENT
		value INTEGER UNIQUE
	);`
	Select = "SELCET value FROM debug;"
	Insert = "INSERT INTO debug(value) VALUES(?);"
	Update = "UPDATE debug SET value = ? WHERE id = 1;"
	Drop   = "DROP TABLE IF EXISTS debug;"
)
