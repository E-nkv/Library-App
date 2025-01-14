Library App: Project to improve my web-dev skills.

The frontend will only be focused initially on the app's functionality, not the styling and UI stuff . Either with CSR (React) or SSR (Go Htmx Templ)

Specs:

# General Flow:

The normal user can see the list of available books and/or authors, and get them with filters & pagination applied.
Logged-in users can also download books, and see book details.

# Authentication:

The api's behaviour depends upon whether the user is logged in or not, and the role of the user.

A user can only have 1 account, determined by email. If he logs in with provider, then we create the corresponding account such that we can keep track of the users data (preferred genres, book status [saved / collection, currently-reading, already-read], and more.)

The user can either create account and login (with unique email), or sign in with provider(s). This
will give the user a jwt or sessionKey (stored in a cookie) which will store the id of the user. On subsequent reqs, api will determine what to do with this token on the auth middleware.

# Routes

1. /books GET PUBLIC BODY={{Filters, Pagination} as JSON}
2. /books/:bookId GET PROTECTED
3. /books/:bookId PUT PROTECTED ADMIN/MODERATOR
4. /books POST CREATE PROTECTED ADMIN BODY={bookData as JSON}
5. /books/:bookId DELETE PROTECTED ADMIN

6. /auth/login POST PUBLIC BODY={{email, password} as JSON}
7. /auth/login GET PUBLIC WITH_PROVIDER
8. /auth/logout PUT PROTECTED
9. /auth/signup POST PUBLIC BODY={{email, password, nick, preferences} as JSON}
