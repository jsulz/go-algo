package models

type User struct {
	ID    int
	First string
	Last  string
}

var (
	users  []*User
	nextID = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(new_user User) (User, error) {
	new_user.ID = nextID
	nextID++
	users = append(users, &new_user)
	return new_user, nil
}
