package models

import (
	"errors"
	"main/utils"
	"time"

	"github.com/google/uuid"
)

type Space struct {
	ID              uuid.UUID `json:"id"`
	Name            string    `json:"name"`
	Location        string    `json:"location"`
	CapacityFrom    int       `json:"capacity_from"`
	CapacityTo      int       `json:"capacity_to"`
	PricePerHour    float64   `json:"price_per_hour"`
	Description     string    `json:"description"`
	CreatedAt       time.Time `json:"created_at"`
	CreatedByUserId uuid.UUID `json:"created_by_user_id"`
	ImageUrl        string    `json:"image_url"`
	Status          string    `json:"status"`
}

func (s *Space) Validate() error {
	if err := utils.ValidateName("name", s.Name, true, 1, 255); err != nil {
		return err
	}
	if err := utils.ValidateName("location", s.Location, true, 5, 255); err != nil {
		return err
	}
	if err := validateCapacity(s.CapacityFrom, s.CapacityTo); err != nil {
		return err
	}
	if err := validatePricePerHour(s.PricePerHour); err != nil {
		return err
	}
	if err := utils.ValidateName("description", s.Description, true, 30, 1000); err != nil {
		return err
	}

	return nil
}

func (s *Space) SetDefaultValues() {
	s.ID = uuid.New()
	s.CreatedAt = time.Now()
	s.Status = "available"
}

func validateCapacity(capacityFrom, capacityTo int) error {
	if capacityFrom < 1 || capacityTo < 1 {
		return errors.New("capacity cannot be less than 1")
	}
	if capacityFrom > capacityTo {
		return errors.New("capacity from cannot be greater than capacity to")
	}
	return nil
}

func validatePricePerHour(pricePerHour float64) error {
	if pricePerHour < 0 {
		return errors.New("price per hour cannot be negative")
	}
	return nil
}
