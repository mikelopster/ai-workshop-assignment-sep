package docs

// This file contains minimal, hand-written swagger docs so the project builds
// without depending on the external `swag` CLI during this edit.

import "github.com/swaggo/swag"

func init() {
	swag.Register(swag.Name, &s{})
}

type s struct{}

func (s *s) ReadDoc() string {
	return `{
		"swagger": "2.0",
		"info": {
			"description": "Simple Profile API for the profile UI used in the workshop.",
			"title": "Profile API",
			"version": "1.0"
		},
		"host": "localhost:3000",
		"basePath": "/",
		"paths": {}
	}`
}
