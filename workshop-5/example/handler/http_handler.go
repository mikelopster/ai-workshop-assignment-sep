package handler

import (
	"example.com/mike/usecase"
	"github.com/gofiber/fiber/v2"
)

// @title           User Management API
// @version         1.0
// @description     A simple user management API with registration and retrieval functionality
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:3000
// @BasePath  /

// HTTPHandler handles HTTP requests
type HTTPHandler struct {
	userUsecase usecase.UserUsecase
}

// NewHTTPHandler creates a new HTTP handler
func NewHTTPHandler(userUsecase usecase.UserUsecase) *HTTPHandler {
	return &HTTPHandler{
		userUsecase: userUsecase,
	}
}

// RegisterRoutes sets up all the routes
func (h *HTTPHandler) RegisterRoutes(app *fiber.App) {
	// Health check endpoint
	app.Get("/health", h.HealthCheck)

	// User endpoints
	app.Post("/register", h.Register)
	app.Get("/user/:id", h.GetUser)
	app.Get("/users", h.GetAllUsers)
}

// HealthCheck handles health check requests
// @Summary      Health Check
// @Description  Check if the service is running
// @Tags         health
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "Service is healthy"
// @Router       /health [get]
func (h *HTTPHandler) HealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "ok",
		"message": "Service is healthy",
	})
}

// Register handles user registration
// @Summary      Register a new user
// @Description  Register a new user with first name, last name, phone, and email
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request  body      usecase.RegisterRequest  true  "User registration data"
// @Success      201      {object}  usecase.RegisterResponse  "User registered successfully"
// @Failure      400      {object}  usecase.RegisterResponse  "Invalid request format or validation error"
// @Failure      409      {object}  usecase.RegisterResponse  "Email already registered"
// @Failure      500      {object}  usecase.RegisterResponse  "Internal server error"
// @Router       /register [post]
func (h *HTTPHandler) Register(c *fiber.Ctx) error {
	var req usecase.RegisterRequest

	// Parse request body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(usecase.RegisterResponse{
			Success: false,
			Message: "Invalid request format",
		})
	}

	// Call usecase
	response, err := h.userUsecase.Register(req)
	if err != nil {
		return c.Status(500).JSON(usecase.RegisterResponse{
			Success: false,
			Message: "Internal server error",
		})
	}

	// Return appropriate status code
	if response.Success {
		return c.Status(201).JSON(response)
	}

	// Check if it's a conflict (email already exists)
	if response.Message == "Email already registered" {
		return c.Status(409).JSON(response)
	}

	return c.Status(400).JSON(response)
}

// GetUser handles getting a user by ID
// @Summary      Get user by ID
// @Description  Retrieve a user by their unique ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  usecase.RegisterResponse  "User found"
// @Failure      404  {object}  usecase.RegisterResponse  "User not found"
// @Failure      500  {object}  usecase.RegisterResponse  "Internal server error"
// @Router       /user/{id} [get]
func (h *HTTPHandler) GetUser(c *fiber.Ctx) error {
	userID := c.Params("id")

	response, err := h.userUsecase.GetUser(userID)
	if err != nil {
		return c.Status(500).JSON(usecase.RegisterResponse{
			Success: false,
			Message: "Internal server error",
		})
	}

	if !response.Success {
		return c.Status(404).JSON(response)
	}

	return c.JSON(response)
}

// GetAllUsers handles getting all users
// @Summary      Get all users
// @Description  Retrieve all registered users
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "List of all users"
// @Failure      500  {object}  map[string]interface{}  "Internal server error"
// @Router       /users [get]
func (h *HTTPHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.userUsecase.GetAllUsers()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"users":   users,
		"count":   len(users),
	})
}
