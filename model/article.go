package model

type Article struct {
	ID           uint    `gorm:"primary_key" json:"id"`
	Title        string  `json:"title"`
	PlainContent string  `gorm:"type:longtext" json:"plain_content"`
	HtmlContent  string  `gorm:"type:longtext" json:"html_content"`
	UserID       uint    `json:"-"`
	User         *User   `json:"user"`
	CreatedAt    MyTime  `json:"created_at"`
	UpdatedAt    MyTime  `json:"updated_at"`
	DeletedAt    *MyTime `json:"deleted_at" sql:"index"`
}
