package domain

type User struct {
	username string
}

type Task struct {
	title       string
	body        string
	priority    int
	asignee     User
	author      User
	storypoints int
}
