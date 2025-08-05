package model

type User struct {
	ID    string  `json:"id" gorm:"primaryKey"`
	Name  string  `json:"name"`
	Email string  `json:"email"`
	Posts []Post  `json:"posts" gorm:"foreignKey:AuthorID"`
}

type NewUserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
