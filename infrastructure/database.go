package infrastructure

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"tracker/domain"
)

type PgStorage struct {
	database *gorm.DB
}

func NewPgStorage(connUrl string) (*PgStorage, error) {
	db, err := gorm.Open(postgres.Open(connUrl), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		return nil, err
	}
	return &PgStorage{database: db}, nil
}

func (storage *PgStorage) AutoMigrate() error {
	return storage.database.AutoMigrate(&domain.User{}, &domain.Task{})
}

// Users

func (storage *PgStorage) CreateUser(username string) (domain.User, error) {
	if username == "" {
		return domain.User{}, fmt.Errorf("username cannot be empty")
	}
	newUser := domain.User{
		ID:       uuid.New(),
		Username: username,
	}

	err := storage.database.Create(&newUser).Error

	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return domain.User{}, fmt.Errorf("username already in use")
		}
		return domain.User{}, err
	}
	return newUser, nil
}

func (storage *PgStorage) GetUser(uid uuid.UUID) (domain.User, error) {
	var user domain.User
	err := storage.database.First(&user, "id = ?", uid).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.User{}, fmt.Errorf("cannot find user with uuid: %s ", uid)
		}
		return domain.User{}, err
	}
	return user, nil
}

func (storage *PgStorage) UpdateUser(uid uuid.UUID, username string) (domain.User, error) {
	if username == "" {
		return domain.User{}, fmt.Errorf("username cannot be empty")
	}

	var updatedUser domain.User
	updModelRes := storage.database.Model(&updatedUser).Where("id = ?", uid).Update("username", username)
	err := updModelRes.Error

	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return domain.User{}, fmt.Errorf("username already in use")
		}
		return domain.User{}, err
	}

	affRows := updModelRes.RowsAffected
	if affRows == 0 {
		return domain.User{}, fmt.Errorf("cannot find user with uuid: %s ", uid)
	}

	return domain.User{ID: uid, Username: username}, nil
}

func (storage *PgStorage) DeleteUser(uid uuid.UUID) error {
	deletionRes := storage.database.Delete(&domain.User{}, "id = ?", uid)
	err := deletionRes.Error

	if err != nil {
		return err
	}

	affRows := deletionRes.RowsAffected
	if affRows == 0 {
		return fmt.Errorf("cannot find user with uuid: %s ", uid)
	}
	return nil
}

// Tasks

func (storage *PgStorage) CreateTask(assignee, author uuid.UUID, body, title string, priority, storypoints int) (domain.Task, error) {
	newTask := domain.Task{
		ID:          uuid.New(),
		Title:       title,
		Body:        body,
		Priority:    priority, // Todo: check exception if blank
		Assignee:    assignee, // Todo: check exception if blank
		Author:      author,   // Todo: check exception if blank
		StoryPoints: storypoints,
	}

	err := storage.database.Create(&newTask).Error
	if err != nil {
		return domain.Task{}, err
	}
	return newTask, err
}

func (storage *PgStorage) GetTask(tuid uuid.UUID) (domain.Task, error) {
	var task domain.Task
	err := storage.database.First(&task, "id = ?", tuid).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Task{}, fmt.Errorf("cannot find task with uuid: %s ", tuid)
		}
		return domain.Task{}, err
	}
	return task, nil
}

func (storage *PgStorage) UpdateTask(tuid uuid.UUID, body, title string, priority, storypoints int) (domain.Task, error) {
	var updatedTask domain.Task
	errGet := storage.database.First(&updatedTask, "id = ?", tuid).Error
	if errGet != nil {
		return domain.Task{}, fmt.Errorf("cannot find task with uuid: %s ", tuid)
	}

	updatedTask.Body = body
	updatedTask.Priority = priority
	updatedTask.StoryPoints = storypoints
	updatedTask.Title = title

	errSave := storage.database.Save(&updatedTask).Error
	if errSave != nil {
		return domain.Task{}, errSave
	}

	return updatedTask, nil
}

func (storage *PgStorage) DeleteTask(tuid uuid.UUID) error {
	delTaskRes := storage.database.Delete(&domain.Task{}, "id = ?", tuid)
	err := delTaskRes.Error
	if err != nil {
		return err
	}

	affRows := delTaskRes.RowsAffected
	if affRows == 0 {
		return fmt.Errorf("cannot find task with uuid: %s ", tuid)
	}
	return nil
}
