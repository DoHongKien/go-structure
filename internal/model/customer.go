package model

const CustomerTableName = "customer"

type Customer struct {
	ID       int    `gorm:"primaryKey;column:id"`
	Username string `gorm:"column:username"`
	FullName string `gorm:"column:full_name"`
	Gender   string `gorm:"column:gender"`
	Status   string `gorm:"column:status"`
	Email    string `gorm:"column:email"`
	Phone    string `gorm:"column:phone_number"`
}

// TableName overrides the default table name used by GORM for the Customer model.
func (Customer) TableName() string {
	return CustomerTableName
}
