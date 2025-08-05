package model

type Post struct {
	ID       string    `json:"id" gorm:"primaryKey"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	AuthorID string    `json:"author_id"`
	Author   *User     `json:"author" gorm:"foreignKey:AuthorID"`
	Comments []Comment `json:"comments" gorm:"foreignKey:PostID"`
}

type NewPostInput struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorID string `json:"authorID"`
}
