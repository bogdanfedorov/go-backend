package user

import "errors"

type authorization struct {
	Token string `json:"token"`
}

type User struct {
	login    string
	password string
}

func login(user User) (*authorization, error) {
	for _, u := range userStorage {
		if u.password == user.password {
			return &authorization{Token: ""}, nil
		}
	}

	return nil, errors.New("bad login or password")
}

func createNewUser(newUser User) (user *User, err error) {
	if checkExistUser(newUser.login) {
		return nil, errors.New("user exist")
	}

	userStorage = append(userStorage, newUser)
	return &newUser, nil
}
