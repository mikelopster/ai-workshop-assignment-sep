package entity

import "time"

// User represents a registered user in the domain
type User struct {
	ID              string    `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	MemberID        string    `json:"member_id" example:"LBK000001"`
	FirstName       string    `json:"first_name" example:"John"`
	LastName        string    `json:"last_name" example:"Doe"`
	Phone           string    `json:"phone" example:"+66812345678"`
	Email           string    `json:"email" example:"john.doe@example.com"`
	MembershipLevel string    `json:"membership_level" example:"Gold"`
	Points          int       `json:"points" example:"0"`
	RegisteredAt    time.Time `json:"registered_at" example:"2024-01-01T00:00:00Z"`
}

// NewUser creates a new user with default values
func NewUser(id, memberID, firstName, lastName, phone, email string) *User {
	return &User{
		ID:              id,
		MemberID:        memberID,
		FirstName:       firstName,
		LastName:        lastName,
		Phone:           phone,
		Email:           email,
		MembershipLevel: "Gold", // Default membership level
		Points:          0,      // Start with 0 points
		RegisteredAt:    time.Now(),
	}
}

// GetFullName returns the full name of the user
func (u *User) GetFullName() string {
	return u.FirstName + " " + u.LastName
}
