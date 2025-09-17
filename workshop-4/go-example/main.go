package main

import (
	"sync"

	"github.com/gofiber/fiber/v2"
)

// Profile represents the data shown on the profile UI.
type Profile struct {
	MembershipLevel string `json:"membership_level"`
	MembershipCode  string `json:"membership_code"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Phone           string `json:"phone"`
	Email           string `json:"email"`
	JoinedDate      string `json:"joined_date"` // ISO date preferred
	Points          int    `json:"points"`
}

var (
	mu      sync.RWMutex
	profile = Profile{
		MembershipLevel: "Gold",
		MembershipCode:  "LBK001234",
		FirstName:       "สมชาย",
		LastName:        "ใจดี",
		Phone:           "081-234-5678",
		Email:           "somchai@example.com",
		JoinedDate:      "2023-06-15",
		Points:          15420,
	}
)

// setupApp builds the Fiber app. Separated so tests can reuse it.
func setupApp() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	app.Get("/profile", func(c *fiber.Ctx) error {
		mu.RLock()
		p := profile
		mu.RUnlock()
		return c.Status(fiber.StatusOK).JSON(p)
	})

	// PUT accepts a full or partial profile JSON and updates server-side store.
	app.Put("/profile", func(c *fiber.Ctx) error {
		var in Profile
		if err := c.BodyParser(&in); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid json"})
		}

		mu.Lock()
		// update only non-empty / non-zero fields so clients can PATCH-like behaviour
		if in.MembershipLevel != "" {
			profile.MembershipLevel = in.MembershipLevel
		}
		if in.MembershipCode != "" {
			profile.MembershipCode = in.MembershipCode
		}
		if in.FirstName != "" {
			profile.FirstName = in.FirstName
		}
		if in.LastName != "" {
			profile.LastName = in.LastName
		}
		if in.Phone != "" {
			profile.Phone = in.Phone
		}
		if in.Email != "" {
			profile.Email = in.Email
		}
		if in.JoinedDate != "" {
			profile.JoinedDate = in.JoinedDate
		}
		if in.Points != 0 {
			profile.Points = in.Points
		}
		p := profile
		mu.Unlock()

		return c.Status(fiber.StatusOK).JSON(p)
	})

	// lightweight swagger UI served from CDN and OpenAPI JSON at /swagger/doc.json
	app.Get("/swagger", func(c *fiber.Ctx) error {
		html := `<!doctype html>
<html>
  <head>
	<meta charset="utf-8" />
	<meta name="viewport" content="width=device-width, initial-scale=1" />
	<title>Swagger UI</title>
	<link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@4/swagger-ui.css" />
  </head>
  <body>
	<div id="swagger"></div>
	<script src="https://unpkg.com/swagger-ui-dist@4/swagger-ui-bundle.js"></script>
	<script>
	  window.onload = function() {
		const ui = SwaggerUIBundle({
		  url: '/swagger/doc.json',
		  dom_id: '#swagger'
		});
	  };
	</script>
  </body>
</html>`

		return c.Type("html").SendString(html)
	})

	app.Get("/swagger/doc.json", func(c *fiber.Ctx) error {
		// minimal OpenAPI document describing the two endpoints
		doc := map[string]interface{}{
			"openapi": "3.0.0",
			"info": map[string]interface{}{
				"title":       "Profile API",
				"version":     "1.0",
				"description": "Simple Profile API for the profile UI used in the workshop.",
			},
			"paths": map[string]interface{}{
				"/profile": map[string]interface{}{
					"get": map[string]interface{}{
						"summary": "Get profile",
						"responses": map[string]interface{}{
							"200": map[string]interface{}{
								"description": "OK",
								"content": map[string]interface{}{
									"application/json": map[string]interface{}{
										"schema": map[string]interface{}{
											"$ref": "#/components/schemas/Profile",
										},
									},
								},
							},
						},
					},
					"put": map[string]interface{}{
						"summary": "Update profile",
						"requestBody": map[string]interface{}{
							"required": true,
							"content": map[string]interface{}{
								"application/json": map[string]interface{}{
									"schema": map[string]interface{}{
										"$ref": "#/components/schemas/Profile",
									},
								},
							},
						},
						"responses": map[string]interface{}{
							"200": map[string]interface{}{
								"description": "OK",
								"content": map[string]interface{}{
									"application/json": map[string]interface{}{
										"schema": map[string]interface{}{
											"$ref": "#/components/schemas/Profile",
										},
									},
								},
							},
						},
					},
				},
			},
			"components": map[string]interface{}{
				"schemas": map[string]interface{}{
					"Profile": map[string]interface{}{
						"type": "object",
						"properties": map[string]interface{}{
							"membership_level": map[string]interface{}{"type": "string"},
							"membership_code":  map[string]interface{}{"type": "string"},
							"first_name":       map[string]interface{}{"type": "string"},
							"last_name":        map[string]interface{}{"type": "string"},
							"phone":            map[string]interface{}{"type": "string"},
							"email":            map[string]interface{}{"type": "string"},
							"joined_date":      map[string]interface{}{"type": "string", "format": "date"},
							"points":           map[string]interface{}{"type": "integer", "format": "int32"},
						},
					},
				},
			},
		}

		return c.Status(fiber.StatusOK).JSON(doc)
	})

	return app
}

func main() {
	app := setupApp()
	// listen on :3000
	app.Listen(":3005")
}
