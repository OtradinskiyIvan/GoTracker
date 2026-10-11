package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type Storage interface {
	// tasks
	CreateTask(assignee, author uuid.UUID, body, title string, priority, storypoints int) (Task, error)
	GetTask(tuid uuid.UUID) (Task, error)
	UpdateTask(tuid, assignee, author uuid.UUID, body, title string, priority, storypoints int) (Task, error)
	DeleteTask(tuid uuid.UUID) error
	// users
	CreateUser(username string) (User, error)
	GetUser(uid uuid.UUID) (User, error)
	UpdateUser(uid uuid.UUID, username string) (User, error)
	DeleteUser(uid uuid.UUID) error
}

type MemoryStorage struct {
	users             map[uuid.UUID]User
	tasks             map[uuid.UUID]Task
	reservedUsernames map[string]bool
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		users:             map[uuid.UUID]User{},
		tasks:             map[uuid.UUID]Task{},
		reservedUsernames: map[string]bool{},
	}
}

// tasks
func (storage *MemoryStorage) CreateTask(assignee, author uuid.UUID, body, title string, priority, storypoints int) (Task, error) {
	if _, ok := storage.users[author]; !ok {
		return Task{}, fmt.Errorf("cannot find the author %s", author)
	}
	if _, ok := storage.users[assignee]; !ok {
		return Task{}, fmt.Errorf("cannot find the assignee %s", assignee)
	}
	newTask := Task{
		ID:          uuid.New(),
		Title:       title,
		Body:        body,
		Priority:    priority,
		Assignee:    assignee,
		Author:      author,
		StoryPoints: storypoints,
	}
	storage.tasks[newTask.ID] = newTask
	return newTask, nil
}

func (storage *MemoryStorage) GetTask(tuid uuid.UUID) (Task, error) {
	task, ok := storage.tasks[tuid]
	if !ok {
		return Task{}, fmt.Errorf("task %s not found", tuid)
	}
	return task, nil
}

func (storage *MemoryStorage) UpdateTask(tuid uuid.UUID, body, title string, priority, storypoints int) (Task, error) {
	reqTask, ok := storage.tasks[tuid]
	if !ok {
		return Task{}, fmt.Errorf("task %s not found", tuid)
	}

	updTask := Task{
		ID:          reqTask.ID,
		Title:       title,
		Body:        body,
		Priority:    priority,
		Assignee:    reqTask.Assignee,
		Author:      reqTask.Author,
		StoryPoints: storypoints,
	}
	storage.tasks[tuid] = updTask
	return updTask, nil
}

func (storage *MemoryStorage) DeleteTask(tuid uuid.UUID) error {
	if _, ok := storage.tasks[tuid]; !ok {
		return fmt.Errorf("task %s not found", tuid)
	}
	delete(storage.tasks, tuid)
	return nil
}

// users
func (storage *MemoryStorage) CreateUser(username string) (User, error) {
	if username == "" {
		return User{}, fmt.Errorf("username cannot be empty")
	}
	if storage.reservedUsernames[username] {
		return User{}, fmt.Errorf("username is already in use")
	}
	newUser := User{
		ID:       uuid.New(),
		Username: username,
	}
	storage.users[newUser.ID] = newUser
	storage.reservedUsernames[username] = true
	return newUser, nil
}

func (storage *MemoryStorage) GetUser(uid uuid.UUID) (User, error) {
	user, ok := storage.users[uid]
	if !ok {
		return User{}, fmt.Errorf("user %s not found", uid)
	}
	return user, nil
}

func (storage *MemoryStorage) UpdateUser(uid uuid.UUID, username string) (User, error) {
	if username == "" {
		return User{}, fmt.Errorf("username cannot be empty")
	}
	reqUser, ok := storage.users[uid]
	if !ok {
		return User{}, fmt.Errorf("user %s not found", uid)
	}
	if storage.reservedUsernames[username] && username != reqUser.Username {
		return User{}, fmt.Errorf("username is already in use")
	}

	updUser := User{
		ID:       reqUser.ID,
		Username: username,
	}
	storage.users[uid] = updUser
	delete(storage.reservedUsernames, reqUser.Username)
	storage.reservedUsernames[username] = true
	return updUser, nil
}

func (storage *MemoryStorage) DeleteUser(uid uuid.UUID) error {
	delUser, ok := storage.users[uid]
	if !ok {
		return fmt.Errorf("user %s not found", uid)
	}
	delete(storage.users, uid)
	delete(storage.reservedUsernames, delUser.Username)
	return nil
}
