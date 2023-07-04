package entities

type UserAdmin struct {
	ID       uint   `gorm:"autoIncrement" json:"id"`
	Name     string `json:"name"`
	Role     int32  `json:"role"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
