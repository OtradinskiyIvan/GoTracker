package domain

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	Username string    `gorm:"unique;not null"`
}

type Task struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Title       string    `gorm:"not null"`
	Body        string
	Priority    int
	Assignee    uuid.UUID `gorm:"type:uuid"`
	Author      uuid.UUID `gorm:"type:uuid;not null"`
	StoryPoints int
}
