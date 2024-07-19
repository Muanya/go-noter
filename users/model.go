package users

import (
	"database/sql"
	"fmt"
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

func FormatRowsToUsers(rows *sql.Rows) ([]User, error) {
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

func GetUsersFromRequest(data *Response) (*User, error) {

	var newUser User
	if value, exists := (*data)[FieldNames[1]]; exists {
		newUser.Email = fmt.Sprintf("%v", value)
	} else {
		return nil, fmt.Errorf("email is Invalid or empty")
	}

	if value, exists := (*data)[FieldNames[2]]; exists {
		newUser.Username = fmt.Sprintf("%v", value)
	} else {
		return nil, fmt.Errorf("username is Invalid or empty")
	}

	if value, exists := (*data)[FieldNames[3]]; exists {
		newUser.Firstname = fmt.Sprintf("%v", value)
	} else {
		return nil, fmt.Errorf("first name is Invalid or empty")
	}

	if value, exists := (*data)[FieldNames[4]]; exists {
		newUser.Lastname = fmt.Sprintf("%v", value)
	} else {
		return nil, fmt.Errorf("last name is Invalid or empty")
	}

	return &newUser, nil

}

func GetUserPassword(data *Response) (string, error) {
	value, exists := (*data)["password"]

	if value == "" || !exists {
		return "", fmt.Errorf("password field missing")
	}

	password, _ := value.(string)

	return password, nil
}
