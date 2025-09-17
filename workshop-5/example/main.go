package main

import (
	"log"

	"example.com/mike/handler"
	"example.com/mike/repository"
	"example.com/mike/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func main() {
	// Create a new Fiber instance
	app := fiber.New()

	// Initialize dependencies (Dependency Injection)
	userRepo := repository.NewMemoryUserRepository()
	userUsecase := usecase.NewUserUsecase(userRepo)
	httpHandler := handler.NewHTTPHandler(userUsecase)

	// Register routes
	httpHandler.RegisterRoutes(app)

	// Swagger documentation
	app.Get("/swagger/doc.json", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "application/json")
		return c.JSON(map[string]interface{}{
			"swagger": "2.0",
			"info": map[string]interface{}{
				"title":       "User Management API",
				"version":     "1.0",
				"description": "A simple user management API with registration and retrieval functionality",
				"contact": map[string]interface{}{
					"name":  "API Support",
					"url":   "http://www.swagger.io/support",
					"email": "support@swagger.io",
				},
				"license": map[string]interface{}{
					"name": "MIT",
					"url":  "https://opensource.org/licenses/MIT",
				},
			},
			"host":     "127.0.0.1:3000",
			"basePath": "/",
			"schemes":  []string{"http"},
			"paths": map[string]interface{}{
				"/health": map[string]interface{}{
					"get": map[string]interface{}{
						"summary":     "Health Check",
						"description": "Check if the service is running",
						"tags":        []string{"health"},
						"responses": map[string]interface{}{
							"200": map[string]interface{}{
								"description": "Service is healthy",
								"schema": map[string]interface{}{
									"type": "object",
								},
							},
						},
					},
				},
				"/register": map[string]interface{}{
					"post": map[string]interface{}{
						"summary":     "Register a new user",
						"description": "Register a new user with first name, last name, phone, and email",
						"tags":        []string{"users"},
						"parameters": []map[string]interface{}{
							{
								"name":        "request",
								"in":          "body",
								"required":    true,
								"description": "User registration data",
								"schema": map[string]interface{}{
									"$ref": "#/definitions/RegisterRequest",
								},
							},
						},
						"responses": map[string]interface{}{
							"201": map[string]interface{}{
								"description": "User registered successfully",
								"schema": map[string]interface{}{
									"$ref": "#/definitions/RegisterResponse",
								},
							},
							"400": map[string]interface{}{
								"description": "Invalid request format or validation error",
								"schema": map[string]interface{}{
									"$ref": "#/definitions/RegisterResponse",
								},
							},
							"409": map[string]interface{}{
								"description": "Email already registered",
								"schema": map[string]interface{}{
									"$ref": "#/definitions/RegisterResponse",
								},
							},
							"500": map[string]interface{}{
								"description": "Internal server error",
								"schema": map[string]interface{}{
									"$ref": "#/definitions/RegisterResponse",
								},
							},
						},
					},
				},
				"/user/{id}": map[string]interface{}{
					"get": map[string]interface{}{
						"summary":     "Get user by ID",
						"description": "Retrieve a user by their unique ID",
						"tags":        []string{"users"},
						"parameters": []map[string]interface{}{
							{
								"name":        "id",
								"in":          "path",
								"required":    true,
								"type":        "string",
								"description": "User ID",
							},
						},
						"responses": map[string]interface{}{
							"200": map[string]interface{}{
								"description": "User found",
								"schema": map[string]interface{}{
									"$ref": "#/definitions/RegisterResponse",
								},
							},
							"404": map[string]interface{}{
								"description": "User not found",
								"schema": map[string]interface{}{
									"$ref": "#/definitions/RegisterResponse",
								},
							},
							"500": map[string]interface{}{
								"description": "Internal server error",
								"schema": map[string]interface{}{
									"$ref": "#/definitions/RegisterResponse",
								},
							},
						},
					},
				},
				"/users": map[string]interface{}{
					"get": map[string]interface{}{
						"summary":     "Get all users",
						"description": "Retrieve all registered users",
						"tags":        []string{"users"},
						"responses": map[string]interface{}{
							"200": map[string]interface{}{
								"description": "List of all users",
								"schema": map[string]interface{}{
									"type": "object",
								},
							},
							"500": map[string]interface{}{
								"description": "Internal server error",
								"schema": map[string]interface{}{
									"type": "object",
								},
							},
						},
					},
				},
			},
			"definitions": map[string]interface{}{
				"User": map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"id": map[string]interface{}{
							"type":    "string",
							"example": "550e8400-e29b-41d4-a716-446655440000",
						},
						"member_id": map[string]interface{}{
							"type":    "string",
							"example": "LBK000001",
						},
						"first_name": map[string]interface{}{
							"type":    "string",
							"example": "John",
						},
						"last_name": map[string]interface{}{
							"type":    "string",
							"example": "Doe",
						},
						"phone": map[string]interface{}{
							"type":    "string",
							"example": "+66812345678",
						},
						"email": map[string]interface{}{
							"type":    "string",
							"example": "john.doe@example.com",
						},
						"membership_level": map[string]interface{}{
							"type":    "string",
							"example": "Gold",
						},
						"points": map[string]interface{}{
							"type":    "integer",
							"example": 0,
						},
						"registered_at": map[string]interface{}{
							"type":    "string",
							"example": "2024-01-01T00:00:00Z",
						},
					},
				},
				"RegisterRequest": map[string]interface{}{
					"type":     "object",
					"required": []string{"first_name", "last_name", "phone", "email"},
					"properties": map[string]interface{}{
						"first_name": map[string]interface{}{
							"type":    "string",
							"example": "John",
						},
						"last_name": map[string]interface{}{
							"type":    "string",
							"example": "Doe",
						},
						"phone": map[string]interface{}{
							"type":    "string",
							"example": "+66812345678",
						},
						"email": map[string]interface{}{
							"type":    "string",
							"example": "john.doe@example.com",
						},
					},
				},
				"RegisterResponse": map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"success": map[string]interface{}{
							"type":    "boolean",
							"example": true,
						},
						"message": map[string]interface{}{
							"type":    "string",
							"example": "User registered successfully",
						},
						"user": map[string]interface{}{
							"$ref": "#/definitions/User",
						},
					},
				},
			},
		})
	})
	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL:         "http://127.0.0.1:3000/swagger/doc.json",
		DeepLinking: false,
	}))

	// Start server
	log.Println("Starting server on port 3000...")
	log.Println("Swagger documentation available at: http://localhost:3000/swagger/")
	log.Fatal(app.Listen(":3000"))
}
