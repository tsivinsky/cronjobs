package data

type User struct {
	Model

	Email       *string `json:"email" gorm:"column:email"`
	Avatar      *string `json:"avatar" gorm:"column:avatar"`
	GitHubLogin string  `json:"githubLogin" gorm:"column:github_login"`
}
