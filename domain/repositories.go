package domain

import "github.com/google/uuid"

type Storage interface {
	CreateTask(body, title string, asignee, author User, priority, storypoints int) (Task, error)
	GetTask(tuid uuid.UUID) (Task, error)
	UpdateTask(tuid uuid.UUID, body, title string, asignee, author User, priority, storypoints int) (Task, error)
	DeleteTask(tuid uuid.UUID) error
}

type MemoryStorage struct {
	users map[uuid.UUID]User
	tasks map[uuid.UUID]Task
}

func (storage *MemoryStorage) CreateTask(body, title string, asignee, author User, priority, storypoints int) (Task, error) {
	newTask := Task{
		body:        body,
		title:       title,
		asignee:     asignee,
		author:      author,
		priority:    priority,
		storypoints: storypoints,
	}
	storage.tasks[uuid.New()] = newTask
	return newTask, nil
}

func (storage *MemoryStorage) GetTask(tuid uuid.UUID) (Task, error) {
	return storage.tasks[tuid], nil
}

func (storage *MemoryStorage) UpdateTask(tuid uuid.UUID, body, title string, asignee, author User, priority, storypoints int) (Task, error) {
	updTask := Task{
		body:        body,
		title:       title,
		asignee:     asignee,
		author:      author,
		priority:    priority,
		storypoints: storypoints,
	}
	storage.tasks[tuid] = updTask
	return updTask, nil
}

func (storage *MemoryStorage) DeleteTask(tuid uuid.UUID) error {
	delete(storage.tasks, tuid)
	return nil
}
