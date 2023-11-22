package entities

type User_tab struct {
	ID       uint `gorm:"primaryKey"`
	Username string
	Password string
}
