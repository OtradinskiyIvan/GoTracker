package domain

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID
	Username string
}

type Task struct {
	ID          uuid.UUID
	Title       string
	Body        string
	Priority    int
	Assignee    uuid.UUID
	Author      uuid.UUID
	StoryPoints int
}
