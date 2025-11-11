package model

type User struct {
	ID       int
	Username string
	Email    string
}

func (u *User) TableName() string {
	return "users"
}
