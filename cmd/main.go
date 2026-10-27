package main

import (
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/joho/godotenv"

	"tracker/domain"
	"tracker/infrastructure"
)

func main() {
	_ = godotenv.Load()

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	pg, err := infrastructure.NewPgStorage(dsn)
	if err != nil {
		log.Fatalf("connect: %v", err)
	}
	if err := pg.AutoMigrate(); err != nil {
		log.Fatalf("migrate: %v", err)
	}

	var store domain.Storage = pg

	// Users
	ivan, err := store.CreateUser("ivan")
	fmt.Printf("create ivan: %+v err: %v\n", ivan, err)

	sanya, err := store.CreateUser("sanya")
	fmt.Printf("create sanya: %+v err: %v\n", sanya, err)

	_, err = store.CreateUser("ivan")
	fmt.Printf("create dup ivan err: %v\n", err) // expect "username already in use"

	got, err := store.GetUser(ivan.ID)
	fmt.Printf("get ivan: %+v err: %v\n", got, err)
	fmt.Printf("match: %v\n", got == ivan)

	// note: UpdateUser is BROKEN right now — this will print "not found" even though ivan exists
	renamed, err := store.UpdateUser(ivan.ID, "ivan2")
	fmt.Printf("rename ivan: %+v err: %v\n", renamed, err)

	// re-create ivan — should now succeed if rename actually freed it
	_, err = store.CreateUser("ivan")
	fmt.Printf("reuse 'ivan' err: %v\n", err)

	// Tasks
	t, err := store.CreateTask(sanya.ID, ivan.ID, "body", "write tests", 2, 3)
	fmt.Printf("create task: %+v err: %v\n", t, err)

	gotT, err := store.GetTask(t.ID)
	fmt.Printf("get task: %+v err: %v\n", gotT, err)
	fmt.Printf("task match: %v\n", gotT == t)

	updT, err := store.UpdateTask(t.ID, "body2", "title2", 1, 5)
	fmt.Printf("update task: %+v err: %v\n", updT, err)
	fmt.Printf("author preserved: %v\n", updT.Author == ivan.ID)

	err = store.DeleteTask(t.ID)
	fmt.Printf("delete task err: %v\n", err)

	_, err = store.GetTask(t.ID)
	fmt.Printf("get deleted task err: %v\n", err)

	// cleanup
	_ = store.DeleteUser(sanya.ID)
	_ = store.DeleteUser(ivan.ID)
	_ = uuid.Nil
}
