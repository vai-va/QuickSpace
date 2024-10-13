package database

import (
	"fmt"
	"main/models"
	"time"
)

func CreateSpace(space models.Space) error {
	fmt.Println("Creating space", space)
	query := `
		INSERT INTO spaces (
			id, 
			name, 
			location,
			capacity_from,
			capacity_to,
			price_per_hour,
			description,
			created_at, 
			created_by_user_id,
			image_url,
			status
		) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := DB.Exec(query,
		space.ID,
		space.Name,
		space.Location,
		space.CapacityFrom,
		space.CapacityTo,
		space.PricePerHour,
		space.Description,
		space.CreatedAt,
		space.CreatedByUserId,
		space.ImageUrl,
		space.Status,
	)

	return err
}

func GetAllSpaces() ([]models.Space, error) {
	query := `
		SELECT 
			id, 
			name, 
			location,
			capacity_from,
			capacity_to,
			price_per_hour,
			description,
			created_at, 
			created_by_user_id,
			image_url,
			status
		FROM spaces
	`

	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var (
		spaces     []models.Space
		created_at string
	)

	for rows.Next() {
		var space models.Space
		err := rows.Scan(
			&space.ID,
			&space.Name,
			&space.Location,
			&space.CapacityFrom,
			&space.CapacityTo,
			&space.PricePerHour,
			&space.Description,
			&created_at,
			&space.CreatedByUserId,
			&space.ImageUrl,
			&space.Status,
		)
		if err != nil {
			return nil, err
		}
		space.CreatedAt, err = time.Parse("2006-01-02 15:04:05", created_at)
		if err != nil {
			return nil, err
		}
		spaces = append(spaces, space)
	}

	return spaces, nil
}

func GetSpaceByID(spaceID string) (models.Space, error) {
	query := `
		SELECT 
			id, 
			name, 
			location,
			capacity_from,
			capacity_to,
			price_per_hour,
			description,
			created_at, 
			created_by_user_id,
			image_url,
			status
		FROM spaces
		WHERE id = ?
	`

	var (
		space      models.Space
		created_at string
	)
	err := DB.QueryRow(query, spaceID).Scan(
		&space.ID,
		&space.Name,
		&space.Location,
		&space.CapacityFrom,
		&space.CapacityTo,
		&space.PricePerHour,
		&space.Description,
		&created_at,
		&space.CreatedByUserId,
		&space.ImageUrl,
		&space.Status,
	)
	if err != nil {
		return space, err
	}
	space.CreatedAt, err = time.Parse("2006-01-02 15:04:05", created_at)
	if err != nil {
		return space, err
	}

	return space, nil
}

func DeleteSpace(spaceID string) error {
	query := `
		DELETE FROM spaces
		WHERE id = ?
	`
	_, err := DB.Exec(query, spaceID)
	return err
}

func GetSpacesByUserIDByEventType(userID string, eventTypeID string) ([]models.Space, error) {
	query := `
		SELECT DISTINCT 
			s.id, 
			s.name, 
			s.location,
			s.capacity_from,
			s.capacity_to,
			s.price_per_hour,
			s.description,
			s.created_at, 
			s.created_by_user_id,
			s.image_url,
			s.status
		FROM spaces s
		INNER JOIN space_reservations sr ON s.id = sr.space_id
		WHERE sr.rented_by_id = ? 
		AND sr.event_type_id = ?
		ORDER BY s.created_at;
	`

	rows, err := DB.Query(query, userID, eventTypeID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var (
		spaces     []models.Space
		created_at string
	)

	for rows.Next() {
		var space models.Space
		err := rows.Scan(
			&space.ID,
			&space.Name,
			&space.Location,
			&space.CapacityFrom,
			&space.CapacityTo,
			&space.PricePerHour,
			&space.Description,
			&created_at,
			&space.CreatedByUserId,
			&space.ImageUrl,
			&space.Status,
		)
		if err != nil {
			return nil, err
		}
		space.CreatedAt, err = time.Parse("2006-01-02 15:04:05", created_at)
		if err != nil {
			return nil, err
		}
		spaces = append(spaces, space)
	}

	return spaces, nil
}
