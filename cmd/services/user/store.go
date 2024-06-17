package user

import (
	"Ecom/types"
	"database/sql"
	"fmt"
)

// THIS IS A REPOSITORY
// This is responsible for communication with database and providing methods for
// for different operations on the database like getByEmail
type Store struct {
	db *sql.DB
}

// This is a constructor that creates a new instance of the repository using an 
// already initialized database which is passed in as argument
func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}


func (s *Store) GetUserByEmail(email string)(*types.User, error){

	// Here fetch a user by the provided email and then return a row of relevant 
	// data or error
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil{
		return nil, err
	}

	// create a space in the memory to hold data of type User
	u := new(types.User)

	// Go through the rows and use the data to populate the u
	for rows.Next(){
		u, err = scanRowIntoUser(rows)
		if err != nil{
			return nil, err
		}
	}

	// if u is not populated, then nothing was found
	if u.ID == 0{
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}


func (s *Store) GetUserByID(id int)(*types.User, error){
	// Here fetch a user by the provided email and then return a row of relevant 
	// data or error
	rows, err := s.db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil{
		return nil, err
	}

	// create a space in the memory to hold data of type User
	u := new(types.User)

	// Go through the rows and use the data to populate the u
	for rows.Next(){
		u, err = scanRowIntoUser(rows)
		if err != nil{
			return nil, err
		}
	}

	// if u is not populated, then nothing was found
	if u.ID == 0{
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func (s *Store) CreateUser(user types.User) error{
	_, err := s.db.Exec(
		"INSERT INTO users (firstName, lastName, email, password) VALUES (?, ?, ?, ?)",
		user.FirstName, user.LastName, user.Email, user.Password,
	)

	if err != nil {
		return err
	}
	return nil
}

// it populates the user object with the data coming from database
func scanRowIntoUser(row *sql.Rows)(*types.User, error){
	user := new(types.User)

	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
