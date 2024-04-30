package model

import "time"

type Task struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// constraint:OnDelete:CASCADE で、ユーザ削除時にタスクも削除
	User User `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	// このフィールドが必要なのは、おそらく↑で foreignKey:UserId を設定するため
	UserId uint `json:"user_id" gorm:"not null"`
}

type TaskResponse struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
