package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

type Item struct {
	LockIndex   int
	Key         string
	Flags       int
	Value       string
	CreateIndex int
	ModifyIndex int
}

func getTest(c *fiber.Ctx) error {
	url := "http://46.175.122.190:8500/v1/kv/test"

	client := fiber.Client{
		UserAgent:                "",
		NoDefaultUserAgentHeader: false,
		JSONEncoder:              nil,
		JSONDecoder:              nil,
	}

	agent := client.Get(url)

	if err := agent.Parse(); err != nil {
		panic(err)
	}

	status, body, errors := agent.Bytes()

	if len(errors) != 0 {
		c.Status(fiber.StatusInternalServerError)
		fmt.Println(status)
		c.JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	res := []Item{}

	err := json.Unmarshal(body, &res)

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		fmt.Println(status)
		c.JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	rawDecodedText, err := base64.StdEncoding.DecodeString(res[0].Value)
	if err != nil {
		panic(err)
	}

	var labels map[string]any
	err = json.Unmarshal(rawDecodedText, &labels)

	if err != nil {
		panic(err)
	}

	c.Status(fiber.StatusOK)
	return c.JSON(labels)
}

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		Next:             nil,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	}))

	app.Get("/test", getTest)

	log.Fatal(app.Listen(":4000"))
}
