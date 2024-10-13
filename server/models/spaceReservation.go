package models

import (
	"errors"
	"main/utils"
	"time"

	"github.com/google/uuid"
)

type SpaceReservation struct {
	ID          uuid.UUID `json:"id"`
	RentedById  uuid.UUID `json:"rented_by_id"`
	SpaceId     uuid.UUID `json:"space_id"`
	EventTypeID uuid.UUID `json:"event_type_id"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (sr *SpaceReservation) Validate() error {
	if err := utils.ValidateTimeOrder(sr.StartTime, sr.EndTime); err != nil {
		return err
	}
	return nil
}

func (sr *SpaceReservation) SetDefaultValues() {
	sr.ID = uuid.New()
	sr.Status = "pending"
	sr.CreatedAt = time.Now()
	sr.UpdatedAt = time.Now()
}

// For patching the status of a space reservation
type UpdateStatusRequest struct {
	Status string `json:"status"`
}

func (usr *UpdateStatusRequest) Validate() error {
	if usr.Status != "pending" && usr.Status != "reserved" && usr.Status != "confirmed" && usr.Status != "cancelled" {
		return errors.New("status must be one of: pending, reserved, confirmed, cancelled")
	}
	return nil
}
