package repository

import "example.com/mike/entity"

// UserRepository defines the interface for user data operations
type UserRepository interface {
	// Create creates a new user
	Create(user *entity.User) error

	// GetByID retrieves a user by ID
	GetByID(id string) (*entity.User, error)

	// GetByEmail retrieves a user by email
	GetByEmail(email string) (*entity.User, error)

	// GetAll retrieves all users
	GetAll() ([]*entity.User, error)

	// Update updates an existing user
	Update(user *entity.User) error

	// Delete deletes a user by ID
	Delete(id string) error
}
