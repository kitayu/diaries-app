CREATE TABLE IF NOT EXISTS diaries (
	id SERIAL PRIMARY KEY,
	title VARCHAR(128) NOT NULL,
	description TEXT NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);