package domain

import "github.com/google/uuid"

type User struct {
	id       uuid.UUID
	username string
}

type Task struct {
	id          uuid.UUID
	title       string
	body        string
	priority    int
	asignee     uuid.UUID
	author      uuid.UUID
	storypoints int
}
