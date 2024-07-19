package users

// Response is a shortcut for map[string]any
type Response map[string]any

// the order should be the order in the User struct
var FieldNames = []string{"id", "email", "user_name", "first_name", "last_name"}
