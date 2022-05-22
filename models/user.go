package models

type User struct{
	Base
	Username string `json:"username"`
	Age int64 `json:"age"`
}