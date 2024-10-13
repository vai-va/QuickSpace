package database

import (
	"main/models"
	"time"
)

func CreateSpaceRating(spaceRating models.SpaceRating) error {
	query := `
		INSERT INTO space_ratings (
			id,
			space_reservation_id,
			rating,
			review,
			created_at,
			renter_reply_id
		)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err := DB.Exec(query,
		spaceRating.ID,
		spaceRating.SpaceReservationID,
		spaceRating.Rating,
		spaceRating.Review,
		spaceRating.CreatedAt,
		spaceRating.RenterReplyID,
	)

	return err
}

func GetSpaceRatingsBySpaceID(spaceID string) ([]models.SpaceRating, error) {
	query := `
		SELECT
			srat.id,
			srat.space_reservation_id,
			srat.rating,
			srat.review,
			srat.created_at,
			srat.renter_reply_id
		FROM space_ratings srat
		JOIN space_reservations sres ON sres.id = srat.space_reservation_id
		WHERE sres.space_id = ?
	`

	rows, err := DB.Query(query, spaceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		spaceRatings []models.SpaceRating
		created_at   string
	)
	for rows.Next() {
		var spaceRating models.SpaceRating
		err := rows.Scan(
			&spaceRating.ID,
			&spaceRating.SpaceReservationID,
			&spaceRating.Rating,
			&spaceRating.Review,
			&created_at,
			&spaceRating.RenterReplyID,
		)
		if err != nil {
			return nil, err
		}

		spaceRating.CreatedAt, err = time.Parse("2006-01-02 15:04:05", created_at)
		if err != nil {
			return nil, err
		}

		spaceRatings = append(spaceRatings, spaceRating)
	}

	return spaceRatings, nil
}

func GetSpaceRatingByID(id string) (models.SpaceRating, error) {
	query := `
		SELECT
			id,
			space_reservation_id,
			rating,
			review,
			created_at,
			renter_reply_id
		FROM space_ratings
		WHERE id = ?
	`

	var (
		spaceRating models.SpaceRating
		createdAt   string
	)
	err := DB.QueryRow(query, id).Scan(
		&spaceRating.ID,
		&spaceRating.SpaceReservationID,
		&spaceRating.Rating,
		&spaceRating.Review,
		&createdAt,
		&spaceRating.RenterReplyID,
	)
	if err != nil {
		return spaceRating, err
	}
	spaceRating.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAt)
	if err != nil {
		return spaceRating, err
	}

	return spaceRating, err
}

func DeleteSpaceRating(id string) error {
	query := `
		DELETE FROM space_ratings
		WHERE id = ?
	`
	_, err := DB.Exec(query, id)
	return err
}

func GetAllSpaceRatings() ([]models.SpaceRating, error) {
	query := `
		SELECT
			id,
			space_reservation_id,
			rating,
			review,
			created_at,
			renter_reply_id
		FROM space_ratings
	`

	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		spaceRatings []models.SpaceRating
		createdAt    string
	)
	for rows.Next() {
		var spaceRating models.SpaceRating
		err := rows.Scan(
			&spaceRating.ID,
			&spaceRating.SpaceReservationID,
			&spaceRating.Rating,
			&spaceRating.Review,
			&createdAt,
			&spaceRating.RenterReplyID,
		)
		if err != nil {
			return nil, err
		}

		spaceRating.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAt)
		if err != nil {
			return nil, err
		}

		spaceRatings = append(spaceRatings, spaceRating)
	}

	return spaceRatings, nil
}
