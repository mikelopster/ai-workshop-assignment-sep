package usecase

import (
	"fmt"

	"example.com/mike/entity"
	"example.com/mike/repository"
	"github.com/google/uuid"
)

// RegisterRequest represents the registration request
type RegisterRequest struct {
	FirstName string `json:"first_name" validate:"required" example:"John"`
	LastName  string `json:"last_name" validate:"required" example:"Doe"`
	Phone     string `json:"phone" validate:"required" example:"+66812345678"`
	Email     string `json:"email" validate:"required,email" example:"john.doe@example.com"`
}

// RegisterResponse represents the registration response
type RegisterResponse struct {
	Success bool         `json:"success" example:"true"`
	Message string       `json:"message" example:"User registered successfully"`
	User    *entity.User `json:"user,omitempty"`
}

// UserUsecase defines the interface for user business operations
type UserUsecase interface {
	// Register registers a new user
	Register(req RegisterRequest) (*RegisterResponse, error)

	// GetUser retrieves a user by ID
	GetUser(id string) (*RegisterResponse, error)

	// GetAllUsers retrieves all users
	GetAllUsers() ([]*entity.User, error)
}

// userUsecase implements the UserUsecase interface
type userUsecase struct {
	userRepo repository.UserRepository
}

// NewUserUsecase creates a new user usecase
func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

// Register registers a new user
func (u *userUsecase) Register(req RegisterRequest) (*RegisterResponse, error) {
	// Validate required fields
	if err := u.validateRegisterRequest(req); err != nil {
		return &RegisterResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	// Check if email already exists
	existingUser, err := u.userRepo.GetByEmail(req.Email)
	if err == nil && existingUser != nil {
		return &RegisterResponse{
			Success: false,
			Message: "Email already registered",
		}, nil
	}

	// Generate unique ID and member ID
	id := uuid.New().String()

	// Get current user count to generate member ID
	users, err := u.userRepo.GetAll()
	if err != nil {
		return &RegisterResponse{
			Success: false,
			Message: "Failed to generate member ID",
		}, err
	}

	memberID := fmt.Sprintf("LBK%06d", len(users)+1)

	// Create new user
	user := entity.NewUser(id, memberID, req.FirstName, req.LastName, req.Phone, req.Email)

	// Save user
	if err := u.userRepo.Create(user); err != nil {
		return &RegisterResponse{
			Success: false,
			Message: "Failed to create user",
		}, err
	}

	return &RegisterResponse{
		Success: true,
		Message: "User registered successfully",
		User:    user,
	}, nil
}

// GetUser retrieves a user by ID
func (u *userUsecase) GetUser(id string) (*RegisterResponse, error) {
	user, err := u.userRepo.GetByID(id)
	if err != nil {
		return &RegisterResponse{
			Success: false,
			Message: "User not found",
		}, err
	}

	return &RegisterResponse{
		Success: true,
		Message: "User found",
		User:    user,
	}, nil
}

// GetAllUsers retrieves all users
func (u *userUsecase) GetAllUsers() ([]*entity.User, error) {
	return u.userRepo.GetAll()
}

// validateRegisterRequest validates the registration request
func (u *userUsecase) validateRegisterRequest(req RegisterRequest) error {
	if req.FirstName == "" {
		return fmt.Errorf("first name is required")
	}
	if req.LastName == "" {
		return fmt.Errorf("last name is required")
	}
	if req.Phone == "" {
		return fmt.Errorf("phone number is required")
	}
	if req.Email == "" {
		return fmt.Errorf("email is required")
	}
	return nil
}
