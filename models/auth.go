package internal

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" gorm:"unique;not null" binding:"required"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (User) TableName() string {
	return "User"
}
