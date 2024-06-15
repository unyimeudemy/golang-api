package types

import "time"

// Here interface is used so that we can just test by creating 
// mock interface.
type UserStore interface {
	GetUserByEmail(email string)(*User, error)
	GetUserByID(id int)(*User, error)
	CreateUser(User) error

}

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

type RegisterUserPayload struct {
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}