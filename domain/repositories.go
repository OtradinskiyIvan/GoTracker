package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type Storage interface {
	//tasks
	CreateTask(asignee, author uuid.UUID, body, title string, priority, storypoints int) (Task, error)
	GetTask(tuid uuid.UUID) (Task, error)
	UpdateTask(tuid, asignee, author uuid.UUID, body, title string, priority, storypoints int) (Task, error)
	DeleteTask(tuid uuid.UUID) error
	// users
	CreateUser(username string) (Task, error)
	GetUser(uuid uuid.UUID) (Task, error)
	UpdateUser(uuid uuid.UUID, username string) (Task, error)
	DeleteUser(uuid uuid.UUID) error
}

type MemoryStorage struct {
	users map[uuid.UUID]User
	tasks map[uuid.UUID]Task
}

// tasks
func (storage *MemoryStorage) CreateTask(asignee, author uuid.UUID, body, title string, priority, storypoints int) (Task, error) {
	_, ok1 := storage.users[author]
	_, ok2 := storage.users[asignee]

	if !(ok1 && ok2) {
		return Task{}, fmt.Errorf("Cannot find asignee and author users")
	}
	newTask := Task{
		id:          uuid.New(),
		body:        body,
		title:       title,
		asignee:     asignee,
		author:      author,
		priority:    priority,
		storypoints: storypoints,
	}
	storage.tasks[newTask.id] = newTask
	return newTask, nil
}

func (storage *MemoryStorage) GetTask(tuid uuid.UUID) (Task, error) {
	_, ok := storage.tasks[tuid]
	if !ok {
		return Task{}, fmt.Errorf("task %s not found", tuid)
	}
	return storage.tasks[tuid], nil
}

func (storage *MemoryStorage) UpdateTask(tuid uuid.UUID, body, title string, priority, storypoints int) (Task, error) {
	reqTask, ok := storage.tasks[tuid]
	if !ok {
		return reqTask, fmt.Errorf("task %s not found", tuid)
	}

	updTask := Task{
		id:          reqTask.id,
		body:        body,
		title:       title,
		asignee:     reqTask.asignee,
		author:      reqTask.author,
		priority:    priority,
		storypoints: storypoints,
	}
	storage.tasks[tuid] = updTask
	return updTask, nil
}

func (storage *MemoryStorage) DeleteTask(tuid uuid.UUID) error {
	_, ok := storage.tasks[tuid]
	if !ok {
		return fmt.Errorf("task %s not found", tuid)
	}
	delete(storage.tasks, tuid)
	return nil
}

// users
func (storage *MemoryStorage) CreateUser(username string) (User, error) {
	newUser := User{
		id:       uuid.New(),
		username: username,
	}
	storage.users[newUser.id] = newUser
	return newUser, nil
}

func (storage *MemoryStorage) GetUser(uuid uuid.UUID) (User, error) {
	_, ok := storage.tasks[uuid]
	if !ok {
		return User{}, fmt.Errorf("task %s not found", uuid)
	}
	return storage.users[uuid], nil
}

func (storage *MemoryStorage) UpdateUser(uuid uuid.UUID, username string) (User, error) {
	reqUser, ok := storage.users[uuid]
	if !ok {
		return reqUser, fmt.Errorf("task %s not found", tuid)
	}

	updUser := User{
		id:       reqUser.id,
		username: username,
	}
	storage.users[uuid] = updUser
	return updUser, nil
}

func (storage *MemoryStorage) DeleteUser(uuid uuid.UUID) error {
	_, ok := storage.tasks[uuid]
	if !ok {
		return fmt.Errorf("task %s not found", uuid)
	}
	delete(storage.tasks, uuid)
	return nil
}
