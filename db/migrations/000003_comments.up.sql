CREATE TABLE IF NOT EXISTS comments (
	book_id BIGINT REFERENCES books(id),
	comment_id INT GENERATED ALWAYS AS IDENTITY,
	user_id BIGINT REFERENCES users(id),
	txt TEXT,
	PRIMARY KEY(book_id, comment_id)
);