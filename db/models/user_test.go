package models

import (
	"errors"
	"fmt"
	"library/db/types"
	"library/errs"
	"testing"
)

func Test_PsqlUserModel_GetUserWithTags(t *testing.T) {
	db, err := InitDB()
	if err != nil {
		panic(err)
	}
	m := PsqlUserModel{db}
	user1, err1 := m.GetUserWithTags(10)
	if err1 != nil {
		fmt.Println(err1)
		t.Errorf("expected nil err, got %v", err)
	} else {
		fmt.Printf("succesfuly got user1!:  \n %+v\n", *user1)
	}
	_, err2 := m.GetUserWithTags(-1)
	if err2 == nil || !errors.Is(err2, errs.ErrNotFound) {
		t.Errorf("expected to get a not found error, got %v", err)
	} else {
		fmt.Println("succesfully got err2 as a notfound!")
	}
}

func Test_PsqlUserModel_GetUsersWithTags(t *testing.T) {
	db, err := InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	m := PsqlUserModel{db}
	users, err := m.GetUsersWithTags(5, 7)
	if err != nil {
		t.Errorf("error in getting users: %v", err)
	}
	t.Log("success! the users are:")
	for _, u := range users {
		t.Logf("%+v", *u)
	}
}

func Test_PsqlUserModel_GetUsers(t *testing.T) {
	db, err := InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	m := PsqlUserModel{db}
	users, err := m.GetUsers(5, 7)
	if err != nil {
		t.Errorf("error in getting users: %v", err)
	}
	t.Log("success! the users are:")
	for _, u := range users {
		t.Logf("%+v", *u)
	}
}

func Test_PsqlUserModel_CreateUser(t *testing.T) {
	db, err := InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	m := PsqlUserModel{db}
	//ensure we start fresh in the db
	m.DB.Exec("DELETE FROM users WHERE full_name IN ('User1', 'User2', 'User3')")
	//tag names in user creation matter not.
	users := []types.UserCreate{
		{
			FullName:    "User1",
			Email:       "asdasd@gmail.com",
			PasswdPlain: "pass1",
			Tags: []types.Tag{
				{ID: 3, Name: "asd"}, {ID: 5, Name: "asd"},
			},
		},
		{
			FullName:    "User2",
			Email:       "zxczxc@gmail.com",
			PasswdPlain: "pass2",
			Tags:        []types.Tag{},
		},
		{
			FullName:    "User3",
			Email:       "vbnvbn@gmail.com",
			PasswdPlain: "pass3",
			Tags: []types.Tag{
				{ID: 300, Name: "asd"}, {ID: 3, Name: "asd"},
			},
		},
	}

	uID, err := m.CreateUser(&users[0])
	if err != nil {
		t.Errorf("expected success for user1, got err %v", err)
	} else {
		t.Log("Succesfully created user id: ", uID)
	}
	uID, err = m.CreateUser(&users[1])
	if err != nil {
		t.Errorf("expected success, got %v", err)
	} else {
		t.Log("Succesfully created user id: ", uID)
	}
	//this test should error in a db where tag with id 12 does not exist. If different db, modify the users id, or the test result might be inconsistent
	uID, err = m.CreateUser(&users[2])
	if err != nil {
		t.Log("succesfuly got an error for usrer2: ", err)
	} else {
		t.Errorf("expected to be an error, got %v", uID)
	}
}
