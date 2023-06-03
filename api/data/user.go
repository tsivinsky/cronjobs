package data

type User struct {
	Model

	Email  *string `json:"email" gorm:"column:email"`
	Avatar *string `json:"avatar" gorm:"column:avatar"`
	Login  string  `json:"login" gorm:"column:login"`
}
