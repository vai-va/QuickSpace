package models

import (
	"errors"
	"regexp"
	"time"

	"main/utils"

	"github.com/google/uuid"
	"github.com/nyaruka/phonenumbers"
	"golang.org/x/crypto/bcrypt"
)

// User represents a user in the system
type User struct {
	ID                uuid.UUID `json:"id"`
	Email             string    `json:"email"`
	Password          string    `json:"password"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	Role              string    `json:"role"`
	PhoneNumber       string    `json:"phone_number"`
	ProfilePictureUrl string    `json:"profile_picture_url"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	CognitoUserId     string    `json:"cognito_user_id"`
}

func (u *User) Validate() error {
	if err := validateEmail(u.Email); err != nil {
		return err
	}
	if err := utils.ValidateName("first_name", u.FirstName, false, 1, 50); err != nil {
		return err
	}
	if err := utils.ValidateName("last_name", u.LastName, false, 1, 50); err != nil {
		return err
	}
	if err := validatePhoneNumber(u.PhoneNumber); err != nil {
		return err
	}
	if err := validatePassword(u.Password); err != nil {
		return err
	}
	return nil
}

func (u *User) SetDefaultValues() error {
	u.ID = uuid.New()
	u.Role = "registered_user"
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	hashedPassword, err := HashPassword(u.Password)
	if err == nil {
		u.Password = hashedPassword
	}
	return err
}

func validateEmail(email string) error {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)

	if !re.MatchString(email) {
		return errors.New("email is invalid")
	}
	if len(email) > 255 {
		return errors.New("email cannot be longer than 255 characters")
	}
	return nil
}

func validatePhoneNumber(phoneNumber string) error {
	_, err := phonenumbers.Parse(phoneNumber, "")
	if err != nil {
		return errors.New("invalid phone number format")
	}

	return nil
}

func validatePassword(password string) error {
	length := len(password)

	switch {
	case length == 0:
		return errors.New("password is required")
	case length < 8:
		return errors.New("password must be at least 8 characters long")
	case length > 255:
		return errors.New("password must not exceed 255 characters")
	default:
		return nil
	}
}

// HashPassword generates a bcrypt hash for the given password.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// VerifyPassword verifies if the given password matches the stored hash.
func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
