URLS & ACTIONS:

# BOOKS
GET /books  -> list of books, with limit pagination, sorting, and filtering by [title, authors, publish date, or tags].
GET /books/{id}
POST /books
PATCH /books/{id}
DELETE /books/{id}
PUT /books/{id}/read -> mark the book as read
PUT /books/{id}/unread -> mark the book as unread
PUT /books/{id}/rank -> rank the book from 1 to 5 stars 
POST /books/{id}/bookmark 
DELETE /books/{id}/unbookmark 
POST /books/id/comments
PUT /books/{bookID}/comments/{commentID} -> edit a comment
DELETE /books/{bookID}/comments/{commentID} 

# USERS
POST /users 
PATCH /users/{id}
PUT /users/{id}/activate URLPARAMS:token
DELETE /users/{id}
GET /users/{id}
GET /users

# AUTHORS
POST /authors
PATCH /authors/{id}
DELETE /authors/{id}
GET /authors/{id}
GET /authors 

# AUTH
POST auth/register
PUT auth/login

GET auth/login/google
GET auth/login/google/callback


DB STUFF:



USERS
id
full_name
email
hashed_password
is_verified


BOOKS
id 
title

BOOK_TAGS
book_id
tag_id

AUTHORS
id
full_name


USER_BOOKSREAD
user_id
read_book_id

USER_BOOKSRANKED
user_id
ranked_book_id
rank 

USER_BOOKSMARKED
user_id
bookmarked_book_id 

USER_BOOKTAGS
user_id
tag_id