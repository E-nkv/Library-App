package main

import "fmt"

type seedBook struct {
	title string
}

var books = []seedBook{
	{"Game of Thrones"},
	{"To Kill a Mockingbird"},
	{"1984"},
	{"Pride and Prejudice"},
	{"The Great Gatsby"},
	{"Moby Dick"},
	{"War and Peace"},
	{"The Catcher in the Rye"},
	{"The Hobbit"},
	{"Fahrenheit 451"},
	{"Brave New World"},
	{"The Lord of the Rings"},
	{"Crime and Punishment"},
	{"The Picture of Dorian Gray"},
	{"Wuthering Heights"},
	{"Jane Eyre"},
	{"The Grapes of Wrath"},
	{"The Chronicles of Narnia"},
	{"The Hitchhiker's Guide to the Galaxy"},
	{"Catch-22"},
	{"Little Women"},
	{"The Bell Jar"},
	{"Animal Farm"},
	{"Lord of the Flies"},
	{"The Alchemist"},
	{"The Road"},
	{"The Fault in Our Stars"},
	{"Dune"},
	{"The Handmaid's Tale"},
	{"The Kite Runner"},
	{"Life of Pi"},
	{"A Tale of Two Cities"},
	{"The Secret Garden"},
	{"Gone with the Wind"},
	{"The Count of Monte Cristo"},
	{"Les Misérables"},
	{"The Old Man and the Sea"},
	{"A Brave New World"},
	{"The Diary of a Young Girl"},
	{"One Hundred Years of Solitude"},
	{"Siddhartha"},
	{"Slaughterhouse-Five"},
	{"The Sound and the Fury"},
	{"Beloved"},
	{"The Road Less Traveled"},
	{"Where the Crawdads Sing"},
	{"Educated"},
}

func (s S) seedBooks(N int) error {
	if N > len(books) {
		return fmt.Errorf("maximum of books is %d", len(books))
	}
	existingTags, err := s.getTagsIds()
	if err != nil {
		return err
	}
	_, err = s.DB.Exec("DELETE FROM books")
	if err != nil {
		return err
	}
	smtpBooks, err := s.DB.Prepare("INSERT INTO books VALUES (DEFAULT,  $1) RETURNING id")
	if err != nil {
		return err
	}
	smtpBookTags, err := s.DB.Prepare("INSERT INTO books_tags VALUES($1, $2)")
	if err != nil {
		return err
	}

	for i := 0; i < N; i++ {
		r := smtpBooks.QueryRow(books[i].title)
		var bookID int64
		if err := r.Scan(&bookID); err != nil {
			fmt.Println("error inserting book: ", i)
			continue
		}
		//insert tags
		tagIdsForBook := selectElsRandomly(existingTags, 1, 4)
		for _, tagId := range tagIdsForBook {
			_, err := smtpBookTags.Exec(bookID, tagId)
			if err != nil {
				fmt.Printf("Error inserting tag %d for book %d\n", tagId, bookID)
			}
		}

	}
	return nil
}

func (s S) getBooksIds() ([]int, error) {
	ids := []int{}

	var u int
	rs, err := s.DB.Query("SELECT id FROM books")
	if err != nil {
		return nil, err
	}
	for i := 0; rs.Next(); i++ {
		err = rs.Scan(&u)
		if err != nil {
			return nil, err
		}
		ids = append(ids, u)
	}

	return ids, nil
}
