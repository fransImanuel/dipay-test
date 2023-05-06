package models

import (
	"time"
)

type Todo struct {
	Todo_Id           int64     `json:"id" gorm:"column:todo_id;primaryKey"`
	Activity_Group_Id int64     `json:"activity_group_id" gorm:"column:activity_group_id"`
	Title             string    `json:"title" gorm:"column:title"`
	Priority          string    `json:"priority" gorm:"column:priority"`
	Is_Active         bool      `json:"is_active" gorm:"column:is_active"`
	Created_At        time.Time `json:"createdAt" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt         time.Time `json:"updatedAt" gorm:"updated_at"`
}

type Test struct {
	Counter int64
}

type PostTodo_Req struct {
	Activity_Group_Id int64  `json:"activity_group_id"  gorm:"column:activity_group_id"`
	Title             string `json:"title" gorm:"column:title"`
	Priority          string `json:"priority" gorm:"column:priority"`
	Is_Active         bool   `json:"is_active" gorm:"column:is_active"`
}

type UpdateTodo_Req struct {
	Title     string `json:"title" binding:"required" gorm:"column:title"`
	Priority  string `json:"priority" gorm:"column:priority"`
	Is_Active bool   `json:"is_active" gorm:"column:is_active"`
}
