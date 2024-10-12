package internal

type Notes struct {
	Id     int    `gorm:"primaryKey"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}
