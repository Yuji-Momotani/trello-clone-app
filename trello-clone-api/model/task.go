package model

import "time"

type Task struct {
	// →task_cardsテーブルを作成
	ID         uint      `json:"id" gorm:"primaryKey"` //intのprimaryKeyを指定することでAuotIncrementも設定される。（GORM）
	Content    string    `json:"content" gorm:"not null"`
	TaskCardID uint      `json:"task_card_id" gorm:"not null"`
	TaskCard   TaskCard  `json:"task" gorm:"foreignKey:TaskCardID; constraint:OnDelete:CASCADE"`
	SortNo     uint      `json:"sort_no" gorm:"not null"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type TaskResponse struct {
	ID         uint   `json:"id"`
	Contetn    string `json:"contetn"`
	TaskCardID uint   `json:"task_card_id"`
	SortNo     uint   `json:"sort_no"`
}
