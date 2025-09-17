package main

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestGetProfile(t *testing.T) {
	app := setupApp()

	req := httptest.NewRequest("GET", "/profile", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	if resp.StatusCode != fiber.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	var p Profile
	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		t.Fatalf("decode failed: %v", err)
	}
	if p.Email == "" || p.FirstName == "" {
		t.Fatalf("unexpected empty profile fields: %+v", p)
	}
}

func TestPutProfileUpdatesFields(t *testing.T) {
	app := setupApp()

	update := map[string]interface{}{
		"first_name": "สมชาย-updated",
		"points":     20000,
	}
	b, _ := json.Marshal(update)
	req := httptest.NewRequest("PUT", "/profile", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	if resp.StatusCode != fiber.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	var p Profile
	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		t.Fatalf("decode failed: %v", err)
	}
	if p.FirstName != "สมชาย-updated" {
		t.Fatalf("first name not updated, got %q", p.FirstName)
	}
	if p.Points != 20000 {
		t.Fatalf("points not updated, got %d", p.Points)
	}
}
