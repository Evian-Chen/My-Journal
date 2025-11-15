package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

const (
	StatusPending    = "pending"
	StatusInProgress = "inProgress"
	StatusDone       = "done"
)

type StatusMeta struct {
	Label string `json:"label"`
	Color string `json:"color"`
}

var StatusMap = map[string]StatusMeta{
	StatusPending: {
		Label: "Pending",
		Color: "#3498db",
	},
	StatusInProgress: {
		Label: "In Progress",
		Color: "#f1c40f",
	},
	StatusDone: {
		Label: "Done",
		Color: "#2ecc71",
	},
}

type Todo struct {
	TodoId  int    `json:"todoId"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  string `json:"status"`
	Color   string `json:"color"`
	Date    string `json:"date"` // 2024-01-02
}

func main() {
	fmt.Println("Hello world again!")
	app := fiber.New()

	todos := map[int]Todo{}
	var todoId int = 1

	// 回傳全部
	app.Get("/api/todo", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"data": todos})
	})

	// 新增一筆 todo
	app.Post("/api/todo", func(c *fiber.Ctx) error {
		todo := new(Todo)
		if err := c.BodyParser(todo); err != nil {
			log.Fatal(err)
			return err
		}

		if todo.Title == "" {
			return c.Status(400).JSON(fiber.Map{"message": "Title is required."})
		}

		if todo.Status == "" {
			todo.Status = StatusMap[StatusPending].Label
			todo.Color = StatusMap[StatusPending].Color
		} else {
			todo.Color = StatusMap[todo.Status].Color
		}
		todo.Date = time.Now().Format("2006-01-02") // 2009-11-10
		todo.TodoId = todoId
		todos[todoId] = *todo
		todoId++
		return c.Status(201).JSON(fiber.Map{"message": "create todo successfully!", "data": (todo)})
	})

	// 更新一筆 todo
	app.Patch("/api/todo/:todoId", func(c *fiber.Ctx) error {
		todoId := c.Params("todoId")
		id, _ := strconv.Atoi(todoId)

		_, exist := todos[id]
		if !exist {
			return c.Status(404).JSON(fiber.Map{"error": "index not found"})
		}

		newTodo := new(Todo)
		if err := c.BodyParser(newTodo); err != nil {
			log.Fatal(err)
			return err
		}

		todos[id] = *newTodo
		return c.Status(201).JSON(fiber.Map{"message": "create todo successfully!", "data": (todos[id])})
	})

	// 刪除一筆 todo

	log.Fatal(app.Listen(":3000"))
}
