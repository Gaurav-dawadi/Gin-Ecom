package models

type User struct{
	Base
	Username string `json:"username"`
}