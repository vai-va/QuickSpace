package database

import "main/models"

func CreateEventType(eventType models.EventType) error {
	query := `
		INSERT INTO event_types (
			id,
			name,
			description,
			icon_url
		)
		VALUES (?, ?, ?, ?)
	`
	_, err := DB.Exec(query,
		eventType.ID,
		eventType.Name,
		eventType.Description,
		eventType.IconUrl,
	)

	return err
}

func GetAllEventTypes() ([]models.EventType, error) {
	query := `
		SELECT
			id,
			name,
			description,
			icon_url
		FROM event_types
	`

	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var eventTypes []models.EventType
	for rows.Next() {
		var eventType models.EventType
		err := rows.Scan(
			&eventType.ID,
			&eventType.Name,
			&eventType.Description,
			&eventType.IconUrl,
		)
		if err != nil {
			return nil, err
		}

		eventTypes = append(eventTypes, eventType)
	}

	return eventTypes, nil
}

func GetEventTypeByID(id string) (models.EventType, error) {
	query := `
		SELECT
			id,
			name,
			description,
			icon_url
		FROM event_types
		WHERE id = ?
	`

	var eventType models.EventType
	err := DB.QueryRow(query, id).Scan(
		&eventType.ID,
		&eventType.Name,
		&eventType.Description,
		&eventType.IconUrl,
	)
	if err != nil {
		return eventType, err
	}

	return eventType, nil
}

func DeleteEventType(id string) error {
	query := `
		DELETE FROM event_types
		WHERE id = ?
	`
	_, err := DB.Exec(query, id)
	return err
}

func PutEventTypeByID(id string, eventType models.EventType) error {
	query := `
		UPDATE event_types
		SET
			name = ?,
			description = ?,
			icon_url = ?
		WHERE id = ?
	`
	_, err := DB.Exec(query,
		eventType.Name,
		eventType.Description,
		eventType.IconUrl,
		id,
	)
	return err
}
