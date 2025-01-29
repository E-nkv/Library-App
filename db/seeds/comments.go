package main

import "fmt"

var comments = []string{
	"Very good, recommend it.",
	"Terrible, the worst.",
	"Wuba luba.",
	"An absolute masterpiece!",
	"Not worth the time.",
	"Could be better.",
	"Exceeded my expectations!",
	"I wouldn't recommend it.",
	"Simply amazing!",
	"Disappointing experience.",
	"Fantastic read!",
	"Too long and boring.",
	"A great addition to my collection.",
	"Meh, it was okay.",
	"Highly informative!",
	"Not my cup of tea.",
	"Would read again!",
	"The plot twist was unexpected!",
	"A bit cliché but enjoyable.",
	"Captivating from start to finish!",
	"I couldn't put it down!",
}

func (s S) seedComments(N int) error {
	usersIds, err := s.getUsersIds()
	if err != nil {
		return err
	}
	booksIds, err := s.getBooksIds()
	if err != nil {
		return err
	}
	if _, err := s.DB.Exec("DELETE FROM comments"); err != nil {
		return err
	}

	for _, bookId := range booksIds {
		commentingUsersIdsForBook := selectElsRandomly(usersIds, 0, N)
		commentsTexts := selectElsRandomly(comments, len(commentingUsersIdsForBook), len(commentingUsersIdsForBook))
		if len(commentingUsersIdsForBook) != len(commentsTexts) {
			panic("ke loco pinga")
		}
		for i := range len(commentsTexts) {
			if _, err := s.DB.Exec("INSERT INTO comments (book_id, user_id, txt) VALUES ($1, $2, $3)", bookId, commentingUsersIdsForBook[i], commentsTexts[i]); err != nil {
				fmt.Printf("Error inserting comment for book %d, with user %d\n", bookId, commentingUsersIdsForBook[i])
			}

		}

	}
	return nil
}
