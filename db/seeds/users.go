package main

import (
	"fmt"
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
		hp, err := bcrypt.GenerateFromPassword([]byte(u.pass), bcrypt.DefaultCost)
		if err != nil {
			fmt.Printf("failed creating pass hash for the user number %d\n", i)
			continue
		}

		r := smtp.QueryRow(u.full_name, hp, u.email)
		var user_id int64
		err = r.Scan(&user_id)
		if err != nil {
			fmt.Printf("failed inserting the user number %d. Error is: %v\n", i, err.Error())
			continue
		}

		tagIdsForUser := genTagsRandomly(existingTags, 0, 10)
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

	}

	return nil
}
