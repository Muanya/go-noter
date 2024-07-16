package models

// id
// username
// password
// email
// created_at
// firstname
// lastname

type Users struct {
	id        int    `json:"id"`
	email     string `json:"email"`
	username  string `json:"user_name"`
	firstname string `json:"first_name"`
	lastname  string `json:"last_name"`
}

// constructor
func New() *Users {
	return &Users{}
}

// getters
func (u *Users) GetId() int {
	return u.id
}

func (u *Users) GetUserName() string {
	return u.username
}

func (u *Users) GetEmail() string {
	return u.email
}

func (u *Users) GetFirstName() string {
	return u.firstname
}

func (u *Users) GetLastName() string {
	return u.lastname
}

// setters for only first name and lastname
func (u *Users) SetFirstName(fName string) {
	u.firstname = fName
}

func (u *Users) SetLastName(lName string) {
	u.lastname = lName
}
