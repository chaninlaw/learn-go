package fiber

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

type User struct {
	Email    string
	Password string
}

type Views interface {
}

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var memberUser = User{
	Email:    "acme@acme.com",
	Password: "acme",
}

var books []Book

func Fiber() {
	if err := godotenv.Load(".env.local"); err != nil {
		fmt.Println("Error loading .env file")
	}

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	books = append(books, Book{ID: 1, Title: "Chanin", Author: "Chanin"})
	books = append(books, Book{ID: 2, Title: "Ninja", Author: "Chanin"})

	config := cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "*",
	})

	auth := jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	})

	app.Use(config, auth, checkMiddleWare)

	configGroup := app.Group("/config")
	configGroup.Use(superadminMiddleware)

	app.Post("/login", login)
	app.Get("/", handleHTML)
	app.Get("/config", getEnv)
	app.Get("/books", getBooks)
	app.Get("/books/:id", getBook)
	app.Post("/books", createBook)
	app.Put("/books/:id", updateBook)
	app.Delete("/books/:id", deleteBook)
	app.Post("/upload", uploadFile)

	app.Listen(":8080")
}
