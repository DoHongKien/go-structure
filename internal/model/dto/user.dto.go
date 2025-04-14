package dto

type UserDTO struct {
	ID       int    `mapstruct:"json:id"`
	Username string `mapstruct:"json:username"`
	Email    string `mapstruct:"json:email"`
	Phone    string `mapstruct:"json:phone"`
	Status   string `mapstruct:"json:status"`
}
