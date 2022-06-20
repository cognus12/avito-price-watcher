package app

const schema = `
	PRAGMA foreign_keys=on;
	
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER UNIQUE NOT NULL PRIMARY KEY AUTOINCREMENT,
		email TEXT UNIQUE
	);

	CREATE TABLE IF NOT EXISTS links (
		id INTEGER UNIQUE NOT NULL PRIMARY KEY AUTOINCREMENT, 
		url TEXT UNIQUE

	);

	CREATE TABLE IF NOT EXISTS subscriptions (
		link_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY (link_id) REFERENCES links(id),
		FOREIGN KEY (user_id) REFERENCES users(id)	
		PRIMARY KEY (link_id, user_id)
	);
`
