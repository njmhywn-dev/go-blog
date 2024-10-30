package model

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserId    string    `json:"user_id" gorm:"column:user_id;unique"`
	Password  string    `json:"-" gorm:"column:password;"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
}
