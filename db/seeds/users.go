package main

import (
	"fmt"
	"math/rand/v2"
	"strings"

	faker "github.com/go-faker/faker/v4"
	"golang.org/x/crypto/bcrypt"
)

type seedUser struct {
	full_name string
	pass      string
	email     string
}

func slicesAt(sc []string, i int) string {
	return sc[len(sc)+i]
}

func genEmailFromName(n string) string {
	sc := strings.Split(n, " ")
	em := strings.ToLower(string(slicesAt(sc, -2)[0]) + slicesAt(sc, -1))
	prov := slicesAt(strings.Split(faker.Email(), "."), -1)
	return em + "." + prov
}
func genSeedUsers(N int, us *[]seedUser) {
	for i := 0; i < N; i++ {
		n := faker.Name()
		em := genEmailFromName(n)
		spl := strings.Split(em, ".")
		u := seedUser{
			full_name: n,
			pass:      strings.Join(spl[:len(spl)-1], ""),
			email:     em,
		}

		(*us)[i] = u
	}

}

func (s S) seedUsers(N int) error {

	existingTags, err := s.getTagsIds()
	if err != nil {
		return err
	}

	smtp, err := s.DB.Prepare(`INSERT INTO users (full_name, hash_pass, email) VALUES($1, $2, $3) RETURNING id`)
	if err != nil {
		return err
	}

	users := make([]seedUser, N)
	_, err = s.DB.Exec("DELETE FROM users_tags")
	if err != nil {
		return err
	}
	_, err = s.DB.Exec("DELETE FROM users")
	if err != nil {
		return err
	}

	genSeedUsers(N, &users)

	for i, u := range users {

		//hash pass
		hp, err := bcrypt.GenerateFromPassword([]byte(u.pass), bcrypt.DefaultCost)
		if err != nil {
			fmt.Printf("failed creating pass hash for the user number %d\n", i)
			continue
		}

		//insert the user and return the id
		r := smtp.QueryRow(u.full_name, hp, u.email)
		var user_id int64
		err = r.Scan(&user_id)
		if err != nil {
			fmt.Printf("failed inserting the user number %d. Error is: %v\n", i, err.Error())
			continue
		}

		//select tags randomly and add them to the users_tags table
		tagIdsForUser := selectElsRandomly(existingTags, 0, 10)
		smtpUsersTags, err := s.DB.Prepare("INSERT INTO users_tags VALUES($1, $2)")
		if err != nil {
			return err
		}
		for _, tagId := range tagIdsForUser {
			_, err := smtpUsersTags.Exec(user_id, tagId)
			if err != nil {
				fmt.Printf("error inserting the tag with id %d for the user %d\n", tagId, user_id)
				fmt.Println(err)
				continue
			}
		}

		booksIds, err := s.getBooksIds()
		if err != nil {
			return err
		}

		//select random books to add to the users_booksRead, users_booksBookmarked, and users_booksRanked tables.
		readBooksIdsForUser := selectElsRandomly(booksIds, 10, 15)
		bookmarkedBooksIdsForUser := selectElsRandomly(booksIds, 4, 8)
		rankedBooksIdsForUser := selectElsRandomly(readBooksIdsForUser, len(readBooksIdsForUser)/2, len(readBooksIdsForUser)*3/4)
		actualRanksForUser := make([]int, len(rankedBooksIdsForUser))

		for i := 0; i < len(rankedBooksIdsForUser); i++ {
			actualRanksForUser[i] = rand.IntN(5) + 1
		}

		readBookSmtp, err := s.DB.Prepare("INSERT INTO users_booksRead VALUES($1, $2)")
		if err != nil {
			return err
		}
		bookmarkedBookSmtp, err := s.DB.Prepare("INSERT INTO users_booksBookmarked VALUES($1, $2)")
		if err != nil {
			return err
		}
		rankedBookSmtp, err := s.DB.Prepare("INSERT INTO users_booksRanked VALUES($1, $2, $3)")
		if err != nil {
			return err
		}

		for _, readBookId := range readBooksIdsForUser {
			if _, err := readBookSmtp.Exec(user_id, readBookId); err != nil {
				fmt.Printf("err inserting readbookid %d for userid %d\n", readBookId, user_id)
			}
		}
		for _, bookmarkedBookId := range bookmarkedBooksIdsForUser {
			if _, err := bookmarkedBookSmtp.Exec(user_id, bookmarkedBookId); err != nil {
				fmt.Printf("err inserting readbookid %d for userid %d\n", bookmarkedBookId, user_id)
			}
		}
		for i, rankedBookId := range rankedBooksIdsForUser {
			if _, err := rankedBookSmtp.Exec(user_id, rankedBookId, actualRanksForUser[i]); err != nil {
				fmt.Printf("err inserting readbookid %d for userid %d\n", rankedBookId, user_id)
			}
		}

	}

	return nil
}

func (s S) getUsersIds() ([]int, error) {
	var out []int
	rs, err := s.DB.Query("SELECT id FROM users")
	if err != nil {
		return nil, err
	}
	var id int
	for rs.Next() {
		err := rs.Scan(&id)
		if err != nil {
			return nil, err
		}
		out = append(out, id)
	}
	return out, nil
}
