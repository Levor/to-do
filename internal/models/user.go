package models

type User struct {
	ID        int64  `json:"id"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	FirstName string `json:"first_name"`
}
