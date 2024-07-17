package users

// id
// username
// password
// email
// created_at
// firstname
// lastname

type Users struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"user_name"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
}

// constructor
func New() *Users {
	return &Users{}
}
