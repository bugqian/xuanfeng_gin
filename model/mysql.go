package model

type User struct {
	AutoId    int64 `gorm:"<-:false"`
	Id        string
	Name      string
	Age       int64
	CreatedAt int64 `gorm:"<-:create"`
	DeletedAt int64 `gorm:"default:1"`
	UpdatedAt int64 `gorm:"<-:create"`
}
