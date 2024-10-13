package database

import (
	"fmt"
	"main/models"
	"time"
)

// Insert a new user into the database
func CreateUser(user models.User) error {

	fmt.Println("Creating user", user)
	query := `
		INSERT INTO users (
			id, 
			email, 
			password, 
			first_name, 
			last_name, 
			phone_number, 
			profile_picture, 
			cognito_user_id
		) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := DB.Exec(query,
		user.ID,
		user.Email,
		user.Password,
		user.FirstName,
		user.LastName,
		user.PhoneNumber,
		user.ProfilePictureUrl,
		user.CognitoUserId,
	)

	return err
}

// Get all users from the database
func GetAllUsers() ([]models.User, error) {
	query := `
		SELECT 
			id, 
			email, 
			password, 
			first_name, 
			last_name, 
			phone_number, 
			profile_picture, 
			created_at, 
			updated_at, 
			cognito_user_id 
		FROM users
	`

	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		users                  []models.User
		created_at, updated_at string
	)
	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.Password,
			&user.FirstName,
			&user.LastName,
			&user.PhoneNumber,
			&user.ProfilePictureUrl,
			&created_at,
			&updated_at,
			&user.CognitoUserId,
		)
		if err != nil {
			return nil, err
		}
		user.CreatedAt, err = time.Parse("2006-01-02 15:04:05", created_at)
		if err != nil {
			return nil, err
		}

		user.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", updated_at)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func GetUserByID(userID string) (models.User, error) {
	query := `
		SELECT 
			id, 
			email, 
			password, 
			first_name, 
			last_name, 
			phone_number, 
			profile_picture, 
			created_at, 
			updated_at, 
			cognito_user_id 
		FROM users
		WHERE id = ?
	`

	var (
		user                   models.User
		created_at, updated_at string
	)
	err := DB.QueryRow(query, userID).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.PhoneNumber,
		&user.ProfilePictureUrl,
		&created_at,
		&updated_at,
		&user.CognitoUserId,
	)
	if err != nil {
		return user, err
	}

	user.CreatedAt, err = time.Parse("2006-01-02 15:04:05", created_at)
	if err != nil {
		return user, err
	}

	user.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", updated_at)
	if err != nil {
		return user, err
	}

	return user, nil
}

func DeleteUser(userID string) error {
	query := `
		DELETE FROM users
		WHERE id = ?
	`

	_, err := DB.Exec(query, userID)
	return err
}
