package entities

const (
	USER_ADMIN_ROLE_READ  int32 = 1
	USER_ADMIN_ROLE_WRITE int32 = 2
)

type UserAdmin struct {
	ID   uint   `gorm:"autoIncrement" json:"id"`
	Name string `json:"name"`
	Role int32  `json:"role"`
	Email string `json:"email"`
}
