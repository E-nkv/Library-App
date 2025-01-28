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
		if i < 3 {
			fmt.Println(u.pass, u.email)
		}
		*us = append(*us, u)
	}

}
func (s S) seedUsers() error {
	smtp, err := s.DB.Prepare(`INSERT INTO users (full_name, hash_pass, email) VALUES($1, $2, $3)`)
	if err != nil {
		return err
	}

	users := make([]seedUser, 0, 20)
	_, err = s.DB.Exec("DELETE FROM users")
	if err != nil {
		return err
	}
	genSeedUsers(20, &users)
	for i, u := range users {
		hp, err := bcrypt.GenerateFromPassword([]byte(u.pass), bcrypt.DefaultCost)
		if err != nil {
			fmt.Printf("failed creating pass hash for the user number %d\n", i)
			continue
		}

		_, err = smtp.Exec(u.full_name, hp, u.email)
		if err != nil {
			fmt.Printf("failed inserting the user number %d. Error is: %v\n", i, err.Error())
			continue
		}

	}

	return nil
}
