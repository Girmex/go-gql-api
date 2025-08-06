package model

type User struct {
	ID       string `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"uniqueIndex"`
	Password string `json:"password"` // ✅ required for login
	Posts    []Post `json:"posts" gorm:"foreignKey:AuthorID"`
}

type NewUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"` // ✅ for registration
}
