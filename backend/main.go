package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title   string             `json:"title" bson:"title"`
	Content string             `json:"content" bson:"content"`
	Status  string             `json:"status" bson:"status"`
	Color   string             `json:"color" bson:"color"`
	Date    string             `json:"date" bson:"date"` // 2024-01-02
}

// pointer to mongo todo_collection
var todo_collection *mongo.Collection

func main() {
	fmt.Println("APP waking up...")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	MONGOURI := os.Getenv("MONGO_URI")

	// connect to mongo
	clientOpt := options.Client().ApplyURI(MONGOURI)
	client, err := mongo.Connect(context.Background(), clientOpt)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())

	fmt.Println("connected to mongo atlas.")

	todo_collection = client.Database("myJournal").Collection("todos")

	app := fiber.New()

	// 添加中間件來處理不同的 Content-Type
	app.Use(func(c *fiber.Ctx) error {
		// 如果 Content-Type 是 text/plain 但內容是 JSON，修正它
		if c.Get("Content-Type") == "text/plain" {
			body := string(c.Body())
			if len(body) > 0 && (body[0] == '{' || body[0] == '[') {
				c.Set("Content-Type", "application/json")
			}
		}
		return c.Next()
	})

	app.Get("/api/todos", getTodos)
	app.Post("/api/todo", createTodos)
	app.Patch("/api/todos/:id", updateTodos)
	// app.Delete("/api/todos/:id", deleteTodos)

	port := os.Getenv("PORT")
	log.Fatal(app.Listen(":" + port))
}

func getTodos(c *fiber.Ctx) error {
	var todos []Todo

	cursor, err := todo_collection.Find(context.Background(), bson.M{})
	if err != nil {
		return err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo Todo
		if err := cursor.Decode(&todo); err != nil {
			return err
		}
		todos = append(todos, todo)
	}

	return c.JSON(todos)
}

func createTodos(c *fiber.Ctx) error {
	todo := new(Todo)

	// 添加調試信息
	fmt.Printf("Raw request body: %s\n", string(c.Body()))
	fmt.Printf("Content-Type: %s\n", c.Get("Content-Type"))

	if err := c.BodyParser(todo); err != nil {
		fmt.Printf("JSON parsing error: %v\n", err)
		return c.Status(400).JSON(fiber.Map{
			"error":         "Invalid JSON format",
			"details":       err.Error(),
			"received_body": string(c.Body()),
		})
	}

	fmt.Printf("Successfully parsed todo: %+v\n", todo)

	if todo.Title == "" {
		return c.Status(400).JSON(fiber.Map{"message": "Title is required."})
	}
	if todo.Status == "" {
		todo.Status = StatusMap[StatusPending].Label
		todo.Color = StatusMap[StatusPending].Color
	} else {
		todo.Color = StatusMap[todo.Status].Color
	}
	if todo.Content == "" {
		todo.Content = ""
	}

	todo.Date = time.Now().Format("2006-01-02") // 2009-11-10

	insertRes, err := todo_collection.InsertOne(context.Background(), todo)
	if err != nil {
		return err
	}

	todo.ID = insertRes.InsertedID.(primitive.ObjectID)

	return c.Status(201).JSON(fiber.Map{"message": "create todo successfully!", "data": (todo)})
}

func updateTodos(c *fiber.Ctx) error {
	id := c.Params("id")
	todo := new(Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid updated body"})
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid ID"})
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{
		"title":   todo.Title,
		"status":  todo.Status,
		"content": todo.Content,
		"color":   todo.Color,
		"date":    todo.Date,
	}}
	_, err = todo_collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "update failed"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "update todo successfully!"})
}

// API withuot DB
// func main() {
// 	fmt.Println("Hello world again!")
// 	app := fiber.New()

// 	todos := map[int]Todo{}
// 	var todoId int = 1

// 	// 回傳全部
// 	app.Get("/api/todo", func(c *fiber.Ctx) error {
// 		return c.Status(200).JSON(fiber.Map{"data": todos})
// 	})

// 	// 新增一筆 todo
// 	app.Post("/api/todo", func(c *fiber.Ctx) error {
// 		todo := new(Todo)
// 		if err := c.BodyParser(todo); err != nil {
// 			log.Fatal(err)
// 			return err
// 		}

// 		if todo.Title == "" {
// 			return c.Status(400).JSON(fiber.Map{"message": "Title is required."})
// 		}

// 		if todo.Status == "" {
// 			todo.Status = StatusMap[StatusPending].Label
// 			todo.Color = StatusMap[StatusPending].Color
// 		} else {
// 			todo.Color = StatusMap[todo.Status].Color
// 		}
// 		todo.Date = time.Now().Format("2006-01-02") // 2009-11-10
// 		todo.TodoId = todoId
// 		todos[todoId] = *todo
// 		todoId++
// 		return c.Status(201).JSON(fiber.Map{"message": "create todo successfully!", "data": (todo)})
// 	})

// 	// 更新一筆 todo
// 	app.Patch("/api/todo/:todoId", func(c *fiber.Ctx) error {
// 		todoId := c.Params("todoId")
// 		id, _ := strconv.Atoi(todoId)

// 		_, exist := todos[id]
// 		if !exist {
// 			return c.Status(404).JSON(fiber.Map{"error": "index not found"})
// 		}

// 		newTodo := new(Todo)
// 		if err := c.BodyParser(newTodo); err != nil {
// 			log.Fatal(err)
// 			return err
// 		}

// 		todos[id] = *newTodo
// 		return c.Status(201).JSON(fiber.Map{"message": "create todo successfully!", "data": (todos[id])})
// 	})

// 	// 刪除一筆 todo
// 	app.Delete(("api/todo/:todoId"), func(c *fiber.Ctx) error {
// 		todoId := c.Params("todoId")
// 		id, _ := strconv.Atoi(todoId)

// 		// 刪除，不使用database
// 		_, exist := todos[id]
// 		if !exist {
// 			return c.Status(404).JSON(fiber.Map{"error": "todo not found"})
// 		}

// 		delete(todos, id)
// 		return c.Status(200).JSON(fiber.Map{"message": "todo deleted successfully!"})
// 	})

// 	log.Fatal(app.Listen(":3000"))
// }
