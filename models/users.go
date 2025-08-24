package models

import "time"

type User struct {

	UserID   uint      `json:"user_id" gorm:"primaryKey;autoIncrement"`

	Name     string    `json:"name"`

	Email    string    `json:"email"`

	JoinDate time.Time `json:"join_date"`

	Role     string    `json:"role" gorm:"type:user_role;default:'user'"`

}