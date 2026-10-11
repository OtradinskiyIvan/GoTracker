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
	CreateUser(username string) (User, error)
	GetUser(uid uuid.UUID) (User, error)
	UpdateUser(uid uuid.UUID, username string) (User, error)
	DeleteUser(uid uuid.UUID) error
}

type MemoryStorage struct {
	users              map[uuid.UUID]User
	tasks              map[uuid.UUID]Task
	reserved_usernames map[string]bool
}

// tasks
func (storage *MemoryStorage) CreateTask(asignee, author uuid.UUID, body, title string, priority, storypoints int) (Task, error) {
	_, ok1 := storage.users[author]
	_, ok2 := storage.users[asignee]

	if !(ok1) {
		return Task{}, fmt.Errorf("Cannot find the author %s", author)
	}
	if !(ok2) {
		return Task{}, fmt.Errorf("Cannot find the asignee %s", asignee)
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
	if len(username) == 0 {
		return User{}, fmt.Errorf("Username cant be empty")
	}
	_, reserved := storage.reserved_usernames[username]
	if reserved {
		return User{}, fmt.Errorf("This username is already in use")
	}
	newUser := User{
		id:       uuid.New(),
		username: username,
	}
	storage.users[newUser.id] = newUser
	storage.reserved_usernames[username] = true
	return newUser, nil
}

func (storage *MemoryStorage) GetUser(uid uuid.UUID) (User, error) {
	_, ok := storage.users[uid]
	if !ok {
		return User{}, fmt.Errorf("user %s not found", uid)
	}
	return storage.users[uid], nil
}

func (storage *MemoryStorage) UpdateUser(uid uuid.UUID, username string) (User, error) {
	if len(username) == 0 {
		return User{}, fmt.Errorf("Username cannot be empty")
	}
	reqUser, ok := storage.users[uid]
	if !ok {
		return User{}, fmt.Errorf("user %s not found", uid)
	}
	_, reserved := storage.reserved_usernames[username]
	if reserved && (username != reqUser.username) {
		return reqUser, fmt.Errorf("This username is already in use")
	}

	updUser := User{
		id:       reqUser.id,
		username: username,
	}
	storage.users[uid] = updUser
	delete(storage.reserved_usernames, reqUser.username)
	storage.reserved_usernames[username] = true
	return updUser, nil
}

func (storage *MemoryStorage) DeleteUser(uid uuid.UUID) error {
	delUser, ok := storage.users[uid]
	if !ok {
		return fmt.Errorf("user %s not found", uid)
	}
	delete(storage.users, uid)
	delete(storage.reserved_usernames, delUser.username)
	return nil
}
