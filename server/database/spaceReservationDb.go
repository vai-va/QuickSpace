package database

import (
	"main/models"
	"time"
)

func CreateSpaceReservation(spaceReservation models.SpaceReservation) error {
	query := `
		INSERT INTO space_reservations (
			id,
			rented_by_id,
			space_id,
			event_type_id,
			start_time,
			end_time,
			status,
			created_at,
			updated_at
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := DB.Exec(query,
		spaceReservation.ID,
		spaceReservation.RentedById,
		spaceReservation.SpaceId,
		spaceReservation.EventTypeID,
		spaceReservation.StartTime,
		spaceReservation.EndTime,
		spaceReservation.Status,
		spaceReservation.CreatedAt,
		spaceReservation.UpdatedAt,
	)

	return err
}

func GetAllSpaceReservations() ([]models.SpaceReservation, error) {
	query := `
		SELECT
			id,
			rented_by_id,
			space_id,
			event_type_id,
			start_time,
			end_time,
			status,
			created_at,
			updated_at
		FROM space_reservations
	`

	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		spaceReservations                        []models.SpaceReservation
		startTime, endTime, createdAt, updatedAt string
	)
	for rows.Next() {
		var spaceReservation models.SpaceReservation
		err := rows.Scan(
			&spaceReservation.ID,
			&spaceReservation.RentedById,
			&spaceReservation.SpaceId,
			&spaceReservation.EventTypeID,
			&startTime,
			&endTime,
			&spaceReservation.Status,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, err
		}

		spaceReservation.StartTime, err = time.Parse("2006-01-02 15:04:05", startTime)
		if err != nil {
			return nil, err
		}

		spaceReservation.EndTime, err = time.Parse("2006-01-02 15:04:05", endTime)
		if err != nil {
			return nil, err
		}

		spaceReservation.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAt)
		if err != nil {
			return nil, err
		}

		spaceReservation.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", updatedAt)
		if err != nil {
			return nil, err
		}

		spaceReservations = append(spaceReservations, spaceReservation)
	}

	return spaceReservations, nil
}

func GetSpaceReservationsByRentedById(rentedById string) ([]models.SpaceReservation, error) {
	query := `
		SELECT
			id,
			rented_by_id,
			space_id,
			event_type_id,
			start_time,
			end_time,
			status,
			created_at,
			updated_at
		FROM space_reservations
		WHERE rented_by_id = ?
	`

	rows, err := DB.Query(query, rentedById)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		spaceReservations                        []models.SpaceReservation
		startTime, endTime, createdAt, updatedAt string
	)
	for rows.Next() {
		var spaceReservation models.SpaceReservation
		err := rows.Scan(
			&spaceReservation.ID,
			&spaceReservation.RentedById,
			&spaceReservation.SpaceId,
			&spaceReservation.EventTypeID,
			&startTime,
			&endTime,
			&spaceReservation.Status,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, err
		}

		spaceReservation.StartTime, err = time.Parse("2006-01-02 15:04:05", startTime)
		if err != nil {
			return nil, err
		}

		spaceReservation.EndTime, err = time.Parse("2006-01-02 15:04:05", endTime)
		if err != nil {
			return nil, err
		}

		spaceReservation.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAt)
		if err != nil {
			return nil, err
		}

		spaceReservation.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", updatedAt)
		if err != nil {
			return nil, err
		}

		spaceReservations = append(spaceReservations, spaceReservation)
	}

	return spaceReservations, nil
}

func GetSpaceReservationByID(spaceReservationID string) (models.SpaceReservation, error) {
	query := `
		SELECT
			id,
			rented_by_id,
			space_id,
			event_type_id,
			start_time,
			end_time,
			status,
			created_at,
			updated_at
		FROM space_reservations
		WHERE id = ?
	`

	var (
		spaceReservation                         models.SpaceReservation
		startTime, endTime, createdAt, updatedAt string
	)
	err := DB.QueryRow(query, spaceReservationID).Scan(
		&spaceReservation.ID,
		&spaceReservation.RentedById,
		&spaceReservation.SpaceId,
		&spaceReservation.EventTypeID,
		&startTime,
		&endTime,
		&spaceReservation.Status,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return spaceReservation, err
	}

	spaceReservation.StartTime, err = time.Parse("2006-01-02 15:04:05", startTime)
	if err != nil {
		return spaceReservation, err
	}

	spaceReservation.EndTime, err = time.Parse("2006-01-02 15:04:05", endTime)
	if err != nil {
		return spaceReservation, err
	}

	spaceReservation.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAt)
	if err != nil {
		return spaceReservation, err
	}

	spaceReservation.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", updatedAt)
	if err != nil {
		return spaceReservation, err
	}

	return spaceReservation, nil
}

func UpdateSpaceReservationStatus(spaceReservationID string, status string) error {
	query := `
		UPDATE space_reservations
		SET status = ?
		WHERE id = ?
	`
	_, err := DB.Exec(query, status, spaceReservationID)

	return err
}
