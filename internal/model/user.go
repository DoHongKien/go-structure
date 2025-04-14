package model

const USER_TABLE_NAME = "user"

type User struct {
	ID       int    `gorm:"primaryKey;column:id"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Email    string `gorm:"column:email"`
	Phone    string `gorm:"column:phone"`
	Status   string `gorm:"column:status"`
}

func (User) TableName() string {
	return USER_TABLE_NAME
}
