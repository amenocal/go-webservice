package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User // slice of pointers to User; working with pointers helps us to avoid copying the data
	nextID = 1     // compiler will infer the type of the variable
)

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}
