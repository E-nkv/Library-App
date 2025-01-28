package main

import (
	"fmt"
	"math/rand/v2"
	"slices"
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

func (s S) getTagsIds() ([]int, error) {
	ids := make([]int, totalAllowedTags)

	rs, err := s.DB.Query("SELECT id FROM tags")
	if err != nil {
		return nil, err
	}
	for i := 0; rs.Next(); i++ {
		rs.Scan(&ids[i])

	}

	return ids, nil
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
		fmt.Println("id of curr user: ", user_id)

		tagIdsForUser := genTagsRandomly(existingTags)
		if slices.Contains(tagIdsForUser, 0) {
			fmt.Println("viva fidel")
		}
		for _, tagId := range tagIdsForUser {
			s.addTagToUser(user_id, tagId)
		}

	}

	return nil
}

func genTagsRandomly(et []int) []int {
	var idxs []int
	var res []int

	N := len(et)
	amount := rand.IntN(11)
	for range amount {
		var n int
		n = rand.IntN(N)

		for slices.Contains(idxs, n) {

			n = rand.IntN(N)
		}
		idxs = append(idxs, n)
		res = append(res, et[n])

	}

	return res
}

func (s S) addTagToUser(user_id int64, tagId int) {
	res, err := s.DB.Exec("INSERT INTO users_tags VALUES($1, $2)", user_id, tagId)
	if err != nil {
		fmt.Printf("error inserting the tag with id %d for the user %d\n", tagId, user_id)
		fmt.Println(err)
		return
	}
	ra, e := res.RowsAffected()
	if e != nil {
		fmt.Println("error with getting rows affected, ", e)
	} else if ra != 1 {
		fmt.Println("error with rows affected not being 1: ", ra)
	}
}
