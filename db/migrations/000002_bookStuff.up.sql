CREATE TABLE IF NOT EXISTS books_tags (
    book_id BIGINT REFERENCES books(id) ON DELETE CASCADE,
    tag_id INT REFERENCES tags(id) ON DELETE CASCADE
);
CREATE INDEX idx_booksTags_bookId ON books_tags(book_id);

CREATE TABLE IF NOT EXISTS users_booksRead (
    user_id BIGINT REFERENCES  users(id) ON DELETE CASCADE,
    book_id BIGINT REFERENCES books(id) ON DELETE CASCADE
);
CREATE INDEX idx_users_booksRead_userId ON users_booksRead(user_id);

CREATE TABLE IF NOT EXISTS users_booksBookmarked (
    user_id BIGINT REFERENCES  users(id) ON DELETE CASCADE,
    book_id BIGINT REFERENCES books(id) ON DELETE CASCADE
);
CREATE INDEX idx_users_booksBookmarked_userId ON users_booksBookmarked(user_id);

CREATE TABLE IF NOT EXISTS users_booksRanked (
    user_id BIGINT REFERENCES  users(id) ON DELETE CASCADE,
    book_id BIGINT REFERENCES books(id) ON DELETE CASCADE,
    rank SMALLINT
);

CREATE INDEX idx_users_booksRanked_userId ON users_booksBookmarked(user_id);


