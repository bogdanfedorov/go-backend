package user

var userStorage []User

func checkExistUser(login string) (exist bool) {
	for _, user := range userStorage {
		if user.login == login {
			return true
		}
	}

	return false
}
