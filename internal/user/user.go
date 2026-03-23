package user

type User struct{
	Id int64 `json:"id"`
	Username string `json:"username"`
	Age int `json:"age"`
	Email string `json:"email"`
	Password string `json:"-"`
}
type RegisterUser struct {
	Username string `json:"username"`
	Age int `json:"age"`
	Email string `json:"email"`
	Password string `json:"password"`
}