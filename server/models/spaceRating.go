package models

import (
	"errors"
	"main/utils"
	"time"

	"github.com/google/uuid"
)

type SpaceRating struct {
	ID                 uuid.UUID  `json:"id"`
	SpaceReservationID uuid.UUID  `json:"space_reservation_id"`
	Rating             int        `json:"rating"`
	Review             string     `json:"review"`
	CreatedAt          time.Time  `json:"created_at"`
	RenterReplyID      *uuid.UUID `json:"renter_reply_id,omitempty"`
}

func (sr *SpaceRating) Validate() error {
	if err := sr.ValidateRating(sr.Rating); err != nil {
		return err
	}
	if err := utils.ValidateName("review", sr.Review, true, 15, 1000); err != nil {
		return err
	}
	return nil
}

func (sr *SpaceRating) SetDefaultValues() {
	sr.ID = uuid.New()
	sr.CreatedAt = time.Now()
}

func (sr *SpaceRating) ValidateRating(rating int) error {
	if rating < 1 || rating > 10 {
		return errors.New("rating must be between 1 and 5")
	}
	return nil
}
