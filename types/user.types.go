package types

// typedef the db user as the user
// type User repository.User

type UserStore interface {
	GetUserById(id int64)
	GetAllUsers()
	AcceptUserReferal(id int64)
	CreateUser(id int64, payload interface{})
	UpdateUser(id int64, payload interface{})
}
