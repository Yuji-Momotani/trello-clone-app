package model

import "time"

type TaskCard struct {
	// →task_cardsテーブルを作成
	ID        uint      `json:"id" gorm:"primaryKey"` //intのprimaryKeyを指定することでAuotIncrementも設定される。（GORM）
	Title     string    `json:"title" gorm:"not null;"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	User      User      `json:"user" gorm:"foreignKey:UserID; constraint:OnDelete:CASCADE"`
	SortNo    uint      `json:"sort_no" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TaskCardResponse struct {
	ID     uint   `json:"id"`
	Titlle string `json:"title"`
	SortNo uint   `json:"sort_no"`
}

// タスクカードとタスクの親子関係をまとめて返すレスポンス
type TaskCardAndTasksResponse struct {
	ID     uint           `json:"id"`
	Titlle string         `json:"title"`
	SortNo uint           `json:"sort_no"`
	Tasks  []TaskResponse `json:"tasks"`
}
