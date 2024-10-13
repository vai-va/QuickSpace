package models

import (
	"main/utils"

	"github.com/google/uuid"
)

type EventType struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IconUrl     string    `json:"icon_url"`
}

func (e *EventType) Validate() error {
	if err := utils.ValidateName("name", e.Name, true, 2, 100); err != nil {
		return err
	}
	if err := utils.ValidateName("description", e.Description, true, 1, 1000); err != nil {
		return err
	}
	return nil
}

func (e *EventType) SetDefaultValues() {
	e.ID = uuid.New()
}
