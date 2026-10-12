package main

import (
	"fmt"
	"tracker/domain"
)

func main() {
	store := domain.NewMemoryStorage()

	// Users
	ivan, err := store.CreateUser("ivan")
	if err != nil {
		panic(err)
	}
	sanya, err := store.CreateUser("sanya")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Created users: %s (%s), %s (%s)\n",
		ivan.Username, ivan.ID, sanya.Username, sanya.ID)

	// Duplicate username
	_, err = store.CreateUser("ivan")
	fmt.Println("duplicate create err:", err)

	// Tasks
	t, err := store.CreateTask(sanya.ID, ivan.ID, "write tests", "testing", 2, 3)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created task: %+v\n", t)

	// Get
	got, err := store.GetTask(t.ID)
	fmt.Printf("Fetched: %+v (err: %v)\n", got, err)

	// Update
	updated, err := store.UpdateTask(t.ID, "write MORE tests", "testing v2", 1, 5)
	fmt.Printf("Updated: %+v (err: %v)\n", updated, err)

	// Rename user
	renamedAlice, err := store.UpdateUser(ivan.ID, "ivan2")
	fmt.Printf("Renamed: %+v (err: %v)\n", renamedAlice, err)

	// Re-using old name
	_, err = store.CreateUser("ivan")
	fmt.Println("recreate old name err:", err)

	// Delete
	err = store.DeleteTask(t.ID)
	fmt.Println("delete task err:", err)

	_, err = store.GetTask(t.ID)
	fmt.Println("get after delete err:", err)
}
