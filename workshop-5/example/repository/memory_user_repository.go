package repository

import (
	"errors"

	"example.com/mike/entity"
)

// memoryUserRepository implements UserRepository using in-memory storage
type memoryUserRepository struct {
	users map[string]*entity.User
}

// NewMemoryUserRepository creates a new in-memory user repository
func NewMemoryUserRepository() UserRepository {
	return &memoryUserRepository{
		users: make(map[string]*entity.User),
	}
}

// Create creates a new user
func (r *memoryUserRepository) Create(user *entity.User) error {
	if user == nil {
		return errors.New("user cannot be nil")
	}

	if user.ID == "" {
		return errors.New("user ID cannot be empty")
	}

	r.users[user.ID] = user
	return nil
}

// GetByID retrieves a user by ID
func (r *memoryUserRepository) GetByID(id string) (*entity.User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// GetByEmail retrieves a user by email
func (r *memoryUserRepository) GetByEmail(email string) (*entity.User, error) {
	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

// GetAll retrieves all users
func (r *memoryUserRepository) GetAll() ([]*entity.User, error) {
	userList := make([]*entity.User, 0, len(r.users))
	for _, user := range r.users {
		userList = append(userList, user)
	}
	return userList, nil
}

// Update updates an existing user
func (r *memoryUserRepository) Update(user *entity.User) error {
	if user == nil {
		return errors.New("user cannot be nil")
	}

	if user.ID == "" {
		return errors.New("user ID cannot be empty")
	}

	if _, exists := r.users[user.ID]; !exists {
		return errors.New("user not found")
	}

	r.users[user.ID] = user
	return nil
}

// Delete deletes a user by ID
func (r *memoryUserRepository) Delete(id string) error {
	if _, exists := r.users[id]; !exists {
		return errors.New("user not found")
	}

	delete(r.users, id)
	return nil
}
