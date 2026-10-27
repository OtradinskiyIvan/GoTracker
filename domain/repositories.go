package domain

import (
	"github.com/google/uuid"
)

type Storage interface {
	// tasks
	CreateTask(assignee, author uuid.UUID, body, title string, priority, storypoints int) (Task, error)
	GetTask(tuid uuid.UUID) (Task, error)
	UpdateTask(tuid uuid.UUID, body, title string, priority, storypoints int) (Task, error)
	DeleteTask(tuid uuid.UUID) error
	// users
	CreateUser(username string) (User, error)
	GetUser(uid uuid.UUID) (User, error)
	UpdateUser(uid uuid.UUID, username string) (User, error)
	DeleteUser(uid uuid.UUID) error
}
