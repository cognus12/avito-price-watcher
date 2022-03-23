package app

const schema = `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER UNIQUE NOT NULL PRIMARY KEY AUTOINCREMENT,
		email TEXT 
	);

	CREATE TABLE IF NOT EXISTS links (
		id INTEGER UNIQUE NOT NULL PRIMARY KEY AUTOINCREMENT, 
		url TEXT 

	);

	CREATE TABLE IF NOT EXISTS subscriptions (
		link_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY (link_id) REFERENCES links(id),
		FOREIGN KEY (user_id) REFERENCES users(id)
	);
`
