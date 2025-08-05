package model

type Comment struct {
    ID      string `json:"id" gorm:"primaryKey"`
    Content string `json:"content"`
    UserID  string `json:"user_id"`             // foreign key
    User    *User  `json:"user" gorm:"foreignKey:UserID"`  // relation
    PostID  string `json:"post_id"`             // foreign key
    Post    *Post  `json:"post" gorm:"foreignKey:PostID"`  // relation
}


type NewCommentInput struct {
	Content  string `json:"content"`
	PostID   string `json:"postID"`
	AuthorID string `json:"authorID"`
}
