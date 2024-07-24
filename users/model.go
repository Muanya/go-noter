package users

import (
	"fmt"

	"github.com/Muanya/go-noter/db"
	"golang.org/x/crypto/bcrypt"
)

// id
// username
// password
// email
// created_at
// firstname
// lastname

type User struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"user_name"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
}

// constructor
func New() *User {
	return &User{}
}

func (user *User) GetByUsername(username string) error {

	err := db.Conn.QueryRow("SELECT id, username, email, firstname, lastname FROM user WHERE username = ?", username).Scan(&user.Id, &user.Username, &user.Email, &user.Firstname, &user.Lastname)

	if err != nil {
		return err
	}

	return nil

}

func (user *User) GetById(id int) error {

	err := db.Conn.QueryRow("SELECT id, username, email, firstname, lastname FROM user WHERE id = ?", id).Scan(&user.Id, &user.Username, &user.Email, &user.Firstname, &user.Lastname)

	if err != nil {
		return err
	}

	return nil

}

func (user *User) GetFromRequest(data *RegisterRequest) error {
	// todo: add authentication to each field

	user.Email = fmt.Sprintf("%v", (*data).Email)
	user.Username = fmt.Sprintf("%v", (*data).Username)
	user.Firstname = fmt.Sprintf("%v", (*data).Firstname)
	user.Lastname = fmt.Sprintf("%v", (*data).Lastname)

	return nil

}

func GetHashPassword(data *RegisterRequest) ([]byte, error) {
	password := (*data).Password

	if password == "" {
		return nil, fmt.Errorf("password field missing")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password")
	}

	return hashedPassword, nil
}

func CompareHashPassword(password, validPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(validPassword), []byte(password))
	if err != nil {
		fmt.Println("Invalid Password:", err)
		return false
	}

	return true
}

func GetAllUsers() ([]User, error) {

	rows, err := db.Conn.Query("SELECT id, username, email, firstname, lastname FROM user")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var usrs []User

	for rows.Next() {
		var user User

		if err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Firstname, &user.Lastname); err != nil {
			return nil, fmt.Errorf("GetAllUsers %q: %v", user.Username, err)
		}
		usrs = append(usrs, user)
	}

	return usrs, nil

}
