package models

import "time"

type Activities struct {
	Activity_Id int64     `json:"id,omitempty" gorm:"column:activity_id;primaryKey"`
	Title       string    `json:"title,omitempty" gorm:"column:title"`
	Email       string    `json:"email,omitempty" gorm:"column:email"`
	Created_At  time.Time `json:"createdAt,omitempty" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty"`
}

type PostActivity_Req struct {
	Title string `json:"title" binding:"required"`
	Email string `json:"email"`
}
