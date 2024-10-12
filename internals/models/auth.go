package internal

type User struct {
	Id       int    `gorm:"primaryKey" json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (User) TableName() string {
	return "User"
}
